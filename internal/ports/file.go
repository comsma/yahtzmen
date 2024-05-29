package ports

import "mime/multipart"

type FilePort interface {
	UploadFile(fileId string, fileHeader *multipart.FileHeader) error
	RetrieveFileUrl(fileId string) (string, error)
}
