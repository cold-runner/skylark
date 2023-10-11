package sms

import "github.com/cold-runner/skylark/internal/pkg/config"

// 简单工厂模式

func NewSmsClient(conf config.Config) Sms {
	switch conf.ServerConfig().Sms {
	case TENCENT:
		return NewTencentSms(conf.TencentSmsConfig())
	}
	panic("无效的sms服务商")
}
