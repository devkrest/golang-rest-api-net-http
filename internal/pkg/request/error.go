package request

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var msgs []string
		for _, fe := range ve {
			msgs = append(msgs, fe.Field()+" is invalid")
		}
		return errors.New(strings.Join(msgs, ", "))
	}
	return err
}
