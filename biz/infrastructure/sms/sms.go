package sms

import (
	"context"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
)

const (
	TENCENT SmsType = "tencent"
)

var (
	client Sms
)

type SmsType string

func (s SmsType) string() string {
	return string(s)
}

type Sms interface {
	SendToSingle(c context.Context, phone string, paramSet []string) error
	SendToMultiple(c context.Context, phoneSet, paramSet []string) error
}

func Init(c config.Conf) {
	sms := c.GetServer().Sms
	switch sms {
	case TENCENT.string():
		client = newTencentSms(c.GetTencentSms())
	default:
		panic(errors.Errorf("无效的sms类型，支持的类型：tencent。当前传入类型为：%s", c.GetServer().Sms))
	}
}

func SetDefault(s Sms) {
	if s == nil {
		panic(errors.New("sms实例不能为空"))
	}
	client = s
}

func GetClient() Sms {
	if client == nil {
		panic(errors.New("sms实例未初始化，请先调用NewSmsClient初始化"))
	}
	return client
}

func NewSmsClient(conf config.Conf) Sms {
	switch conf.GetServer().Sms {
	case TENCENT.string():
		return newTencentSms(conf.GetTencentSms())
	default:
		panic(errors.Errorf("无效的sms实例类型，支持的类型：tencent。当前传入类型为：%s", conf.GetServer().Sms))
	}
}
