package util

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
	"os"
)

// TODO 新建类型处理
var FILE_TYPE_MAP = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
}

func FileType(file *os.File) (fileType string, err error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}

func FormFileType(fileHeader *multipart.FileHeader) (fileType string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	mime, err := mimetype.DetectReader(file)

	return FILE_TYPE_MAP[mime.String()], nil
}

func IsFileImage(file *os.File) (bool, error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false, err
	}
	if !mime.Is("image/jpeg") && !mime.Is("image/png") {
		return false, nil
	}
	return true, nil
}

func IsFormFileImage(fileHeader *multipart.FileHeader) (bool, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false, err
	}
	if !mime.Is(consts.MIMEImageJPEG) && !mime.Is(consts.MIMEImagePNG) {
		return false, nil
	}
	return true, nil
}
