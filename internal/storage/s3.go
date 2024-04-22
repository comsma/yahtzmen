package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	ProductImagesPath = "product-images"
)

type S3Storage struct {
	bucket string
	cli    *minio.Client
}

func (s *S3Storage) Upload(ctx context.Context, location, key string, file *multipart.FileHeader) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = s.cli.PutObject(
		ctx,
		s.bucket,
		fmt.Sprintf("%s/%s", location, key),
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Storage) Delete(ctx context.Context, path, id string) error {
	// TODO implement me
	panic("implement me")
}

func (s *S3Storage) GetURL(ctx context.Context, location, key string) (string, error) {
	url, err := s.cli.PresignedGetObject(ctx, s.bucket, fmt.Sprintf("%s/%s", location, key), 24*time.Hour, nil)
	if err != nil {
		return "", err

	}
	return url.String(), nil
}

func NewS3Storage(endpoint, accessKey, secret string) *S3Storage {
	client, err := minio.New(
		endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				accessKey,
				secret,
				"",
			),
			Secure: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return &S3Storage{
		bucket: "yahtzmen-storage",
		cli:    client,
	}
}
