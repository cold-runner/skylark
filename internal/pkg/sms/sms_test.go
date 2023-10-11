package sms

import (
	"github.com/cold-runner/skylark/internal/pkg/config"
	"testing"
)

type testSms struct {
	Sms
}

func TestSendSms(t *testing.T) {
	testClient := &testSms{NewTencentSms(&config.tencentSms{
		SecretId:      "",
		SecretKey:     "",
		Region:        "ap-beijing",
		ApplicationId: "",
		SignName:      "",
		TemplateId:    "",
	})}

	err := testClient.Sms.SendToSingle("12345678901", []string{"123456", "5"})
	//err := testClient.Sms.SendToSingle("13348648943", "1234")
	if err != nil {
		t.Errorf("发送短信失败 err: %v", err)
	}
}
