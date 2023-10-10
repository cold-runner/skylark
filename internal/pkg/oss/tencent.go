package oss

import (
	"context"
	"fmt"
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
)

type tencent struct {
	context    context.Context
	client     *cos.Client
	domain     string
	dictionary string
}

func NewTencentOss(opt *option.TencentCos) Oss {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", opt.Bucket, opt.Region))
	c := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  opt.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: opt.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return &tencent{context.Background(), c, opt.Domain, opt.Dictionary}
}

// UploadFormFile 上传文件，请保证文件名唯一
func (t *tencent) UploadFormFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	fileType, err := util.FormFileType(fileHeader)
	if err != nil {
		return "", err
	}

	filePath := t.dictionary + "/" + fileName + "." + fileType
	_, err = t.client.Object.Put(t.context, filePath, file, nil)
	if err != nil {
		return "", err
	}

	return "https://" + t.domain + filePath, nil
}
