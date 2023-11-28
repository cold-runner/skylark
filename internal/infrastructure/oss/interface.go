package oss

import (
	"context"
	"mime/multipart"
)

type Oss interface {
	UploadFormFile(c context.Context, fileHeader *multipart.FileHeader, fileName string) (fileUrl string, err error)
}
