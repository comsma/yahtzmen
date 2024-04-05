package data

import (
	"bytes"
	"errors"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"yahtzmen/internal/storage"
)

var (
	ErrUploadFailed = errors.New("failed to upload file")
)

type File struct {
	URL  string
	UUID string
	Data []byte
}

type FileUploader struct {
	S3 storage.S3
}

// Upload uploads a file to S3
func (f *FileUploader) Upload(file *File) error {

	fr := bytes.NewReader(file.Data)

	_, err := f.S3.Uploader.Upload(
		&s3manager.UploadInput{
			Bucket: &f.S3.Bucket,
			Key:    &file.UUID,
			Body:   fr,
		},
	)
	if err != nil {
		return ErrUploadFailed
	}

	return nil
}

// GetURL retrieves the URL for the file a
func (f *File) GetURL(file *File) error {

}
