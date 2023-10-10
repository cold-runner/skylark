package sms

import (
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 当前API只支持发送中国区域的手机号，也就是+86
type tencentSms struct {
	client        *sms.Client
	applicationId string
	signName      string
	templateId    string
}

func NewTencentSms(opt *option.TencentSms) Sms {
	credential := common.NewCredential(opt.SecretId, opt.SecretKey)
	client, err := sms.NewClient(credential, opt.Region, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	return &tencentSms{
		client:        client,
		applicationId: opt.ApplicationId,
		signName:      opt.SignName,
		templateId:    opt.TemplateId,
	}
}

func (t *tencentSms) buildBasicRequest() *sms.SendSmsRequest {
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(t.applicationId)
	request.SignName = common.StringPtr(t.signName)
	request.TemplateId = common.StringPtr(t.templateId)
	return request
}

func (t *tencentSms) SendToSingle(phone string, paramSet []string) error {
	request := t.buildBasicRequest()
	request.TemplateParamSet = common.StringPtrs(paramSet)
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})

	_, err := t.client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}

func (t *tencentSms) SendToMultiple(phones []string, paramSet []string) error {
	if len(phones) > 200 {
		return errors.New("手机号数量超过200")
	}
	request := t.buildBasicRequest()
	request.TemplateParamSet = common.StringPtrs(paramSet)
	for i := range phones {
		phones[i] = "+86" + phones[i]
	}
	request.PhoneNumberSet = common.StringPtrs(phones)

	_, err := t.client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
