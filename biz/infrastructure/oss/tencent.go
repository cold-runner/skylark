package oss

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/util"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type tencent struct {
	client     *cos.Client
	domain     string
	dictionary string
}

func newTencentOss(opt *config.TencentCos) Oss {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", opt.Bucket, opt.Region))
	c := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  opt.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: opt.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return &tencent{c, opt.Domain, opt.Dictionary}
}

// UploadFormFile 上传文件，请保证文件名唯一
func (t *tencent) UploadFormFile(c context.Context, fileHeader *multipart.FileHeader, fileName string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	fileType, err := util.FormFileType(fileHeader)
	if err != nil {
		return "", err
	}

	filePath := t.dictionary + "/" + fileName + "." + fileType
	_, err = t.client.Object.Put(c, filePath, file, nil)
	if err != nil {
		return "", err
	}

	return "https://" + t.domain + filePath, nil
}

func (t *tencent) UploadFileFromBytes(c context.Context, b []byte, fileName, fileType string) (string, error) {
	filePath := t.dictionary + "/" + fileName + "." + fileType
	file := bytes.NewReader(b)
	_, err := t.client.Object.Put(c, filePath, file, nil)
	if err != nil {
		return "", err
	}
	return "https://" + t.domain + filePath, nil
}
