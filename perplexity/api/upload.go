package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Upload accepts an open file or any reader. The caller owns Reader.
type Upload struct {
	Filename string
	Reader   io.Reader
}

// UploadPath reads a file so retries remain independent of a file's seek position.
func UploadPath(path string) (Upload, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Upload{}, err
	}
	return Upload{Filename: filepath.Base(path), Reader: bytes.NewReader(data)}, nil
}

func Multipart(file Upload, fields map[string]any) ([]byte, string, error) {
	if file.Reader == nil {
		return nil, "", fmt.Errorf("file reader is required")
	}
	name := file.Filename
	if name == "" {
		if f, ok := file.Reader.(interface{ Name() string }); ok {
			name = filepath.Base(f.Name())
		} else {
			name = "upload"
		}
	}
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return nil, "", err
	}
	if _, err = io.Copy(part, file.Reader); err != nil {
		return nil, "", err
	}
	for key, value := range fields {
		if err := writer.WriteField(key, fmt.Sprint(value)); err != nil {
			return nil, "", err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body.Bytes(), writer.FormDataContentType(), nil
}
