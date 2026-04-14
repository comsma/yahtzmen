package storage

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Store struct {
	bucket        string
	publicBaseURL string
	client        *minio.Client
}

func New(endpoint, bucket, publicBaseURL, accessKey, secretKey string) (*Store, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}
	return &Store{
		bucket:        bucket,
		publicBaseURL: strings.TrimRight(publicBaseURL, "/"),
		client:        client,
	}, nil
}

// Upload streams r to the bucket under objectKey.
func (s *Store) Upload(ctx context.Context, objectKey string, r io.Reader, size int64, contentType string) error {
	_, err := s.client.PutObject(ctx, s.bucket, objectKey, r, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// PublicURL returns the public URL for a given object key.
func (s *Store) PublicURL(objectKey string) string {
	return fmt.Sprintf("%s/%s", s.publicBaseURL, objectKey)
}
