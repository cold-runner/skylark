package oss

import (
	"context"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
	"mime/multipart"
)

type OssType string

const (
	TENCENT OssType = "tencent"
)

var (
	o Oss
)

func (o OssType) string() string {
	return string(o)
}

func Init(c config.Conf) {
	oss := c.GetServer().Oss
	switch oss {
	case TENCENT.string():
		o = newTencentOss(c.GetTencentCos())
	default:
		panic(errors.Errorf("无效的对象存储类型，支持的类型：tencent。当前传入类型为：%s", c.GetServer().Oss))
	}
}

type Oss interface {
	UploadFormFile(c context.Context, fileHeader *multipart.FileHeader, fileName string) (fileUrl string, err error)
}
