package data

import (
	"errors"
)

var (
	ErrUploadFailed = errors.New("failed to upload file")
)

type File struct {
	URL  string
	UUID string
	Data []byte
}
