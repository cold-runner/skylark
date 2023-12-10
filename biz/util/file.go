package util

import (
	"encoding/base64"
	"mime/multipart"
	"os"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gabriel-vasile/mimetype"
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

	return FILE_TYPE_MAP[mime.String()], nil
}

func FileTypeFromBs64(bs64 string) (fileType string, err error) {
	data, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return "", err
	}
	mime := mimetype.Detect(data)
	return FILE_TYPE_MAP[mime.String()], nil
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

func IsFormFileImage(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false
	}
	if !mime.Is(consts.MIMEImageJPEG) && !mime.Is(consts.MIMEImagePNG) {
		return false
	}
	return true
}
