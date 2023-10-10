package oss

import (
	"mime/multipart"
)

type Oss interface {
	UploadFormFile(fileHeader *multipart.FileHeader, fileName string) (fileUrl string, err error)
}
