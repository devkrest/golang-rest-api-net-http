package request

import (
	"encoding/json"
	"errors"
	"mime"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"strconv"
)

var validate = validator.New()

// PaginationRequest is a common struct for pagination
type PaginationRequest struct {
	Limit  int `query:"limit" validate:"omitempty,min=1,max=100"`
	Offset int `query:"offset" validate:"omitempty,min=0"`
}

// BindQuery populates a struct from URL query parameters
func BindQuery(r *http.Request, dst any) error {
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("query")
		if tag == "" {
			continue
		}

		val := r.URL.Query().Get(tag)
		if val == "" {
			continue
		}

		fieldV := v.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		switch fieldV.Kind() {
		case reflect.String:
			fieldV.SetString(val)
		case reflect.Int:
			if iVal, err := strconv.Atoi(val); err == nil {
				fieldV.SetInt(int64(iVal))
			}
		case reflect.Int64:
			if iVal, err := strconv.ParseInt(val, 10, 64); err == nil {
				fieldV.SetInt(iVal)
			}
		case reflect.Bool:
			if bVal, err := strconv.ParseBool(val); err == nil {
				fieldV.SetBool(bVal)
			}
		}
	}

	return validate.Struct(dst)
}

func ValidateStruct(v any) error {
	return validate.Struct(v)
}

func BindAndValidate(r *http.Request, dst any) error {
	return Bind(r, dst)
}

func Bind(r *http.Request, dst any) error {

	ct := r.Header.Get("Content-Type")
	mediaType, _, _ := mime.ParseMediaType(ct)

	switch {
	case mediaType == "application/json":
		if err := bindJSON(r, dst); err != nil {
			return err
		}

	case mediaType == "multipart/form-data":
		if err := bindMultipart(r, dst); err != nil {
			return err
		}

	default:
		return errors.New("unsupported content type")
	}

	// STRUCT VALIDATION (JSON + multipart)
	if err := validate.Struct(dst); err != nil {
		return err
	}

	return nil
}

func bindJSON(r *http.Request, dst any) error {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(dst)
}

func bindMultipart(r *http.Request, dst any) error {

	if err := r.ParseMultipartForm(20 << 20); err != nil {
		return err
	}

	// Bind normal fields
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		formTag := f.Tag.Get("form")
		if formTag == "" {
			continue
		}

		val := r.FormValue(formTag)
		if val == "" {
			continue
		}

		v.Field(i).SetString(val)
	}

	// Bind files
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fileTag := f.Tag.Get("file")
		if fileTag == "" {
			continue
		}

		files := r.MultipartForm.File[fileTag]
		if len(files) == 0 {
			continue
		}

		v.Field(i).Set(reflect.ValueOf(files[0]))
	}

	return nil
}
