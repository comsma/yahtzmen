package storage

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	Bucket     string
	Downloader *s3manager.Downloader
	Uploader   *s3manager.Uploader
}

func NewS3(sess *session.Session) S3 {
	return S3{
		Downloader: s3manager.NewDownloader(sess),
		Uploader:   s3manager.NewUploader(sess),
	}
}
