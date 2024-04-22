package storage

import (
	"context"
	"mime/multipart"
)

type Storage interface {
	// Upload uploads file to the storage
	Upload(ctx context.Context, location, key string, file *multipart.FileHeader) error
	// Delete deletes file from the storage
	Delete(ctx context.Context, path, id string) error
	// GetURL retrieves the URL for the file
	GetURL(ctx context.Context, path, id string) (string, error)
}
