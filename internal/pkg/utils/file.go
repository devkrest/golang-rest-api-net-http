package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type UploadResult struct {
	Field    string
	Name     string
	Path     string
	MimeType string
}

// SaveSingle saves an uploaded file and performs security checks
func SaveSingle(
	file *multipart.FileHeader,
	folder string,
	allowedTypes ...string,
) (*UploadResult, error) {

	if file == nil {
		return nil, nil
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// 1. Sniff Content-Type (Security: Don't trust the header)
	buff := make([]byte, 512)
	if _, err := src.Read(buff); err != nil {
		return nil, err
	}
	contentType := http.DetectContentType(buff)
	if _, err := src.Seek(0, 0); err != nil {
		return nil, err
	}

	// 2. Validate Allowed Types (if provided)
	if len(allowedTypes) > 0 {
		allowed := false
		for _, t := range allowedTypes {
			if strings.HasPrefix(contentType, t) {
				allowed = true
				break
			}
		}
		if !allowed {
			return nil, fmt.Errorf("file type %s not allowed", contentType)
		}
	}

	// 3. Generate Secure Unique Filename
	ext := filepath.Ext(file.Filename)
	name := uuid.New().String() + ext
	path := filepath.Join(folder, name)

	if err := os.MkdirAll(folder, 0755); err != nil {
		return nil, err
	}

	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return &UploadResult{
		Name:     name,
		Path:     path,
		MimeType: contentType,
	}, nil
}
