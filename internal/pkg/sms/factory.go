package sms

import "github.com/cold-runner/skylark/internal/pkg/config"

// 简单工厂模式

func NewInstance(conf config.Config) Sms {
	switch conf.ServerConfig().Sms {
	case TENCENT:
		return newTencentSms(conf.TencentSmsConfig())
	}
	panic("无效的sms服务商")
}
