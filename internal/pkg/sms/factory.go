package sms

import "github.com/cold-runner/skylark/internal/pkg/option"

// 简单工厂模式

func NewSmsClient(conf *option.Conf) Sms {
	switch conf.Server.Sms {
	case TENCENT:
		return NewTencentSms(conf.TencentSms)
	}
	panic("无效的sms服务商")
}
