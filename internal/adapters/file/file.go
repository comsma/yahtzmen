package file

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"time"
)

type Adapter struct {
	minio *minio.Client
}

func NewAdapter(endpoint, accessKey, secret string) *Adapter {
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

	return &Adapter{
		minio: client,
	}
}

func (a Adapter) UploadFile(fileId string, fileHeader *multipart.FileHeader) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = a.minio.PutObject(
		context.Background(),
		"yahtzmen-storage",
		fmt.Sprintf("%s/%s", "product-images", fileId),
		src,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (a Adapter) RetrieveFileUrl(fileId string) (string, error) {
	url, err := a.minio.PresignedGetObject(context.Background(), "yahtzmen-storage", fmt.Sprintf("%s/%s", "product-images", fileId), 24*time.Hour, nil)
	if err != nil {
		return "", err

	}
	return url.String(), nil
}
