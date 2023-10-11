package oss

import "github.com/cold-runner/skylark/internal/pkg/config"

// 简单工厂模式

func NewInstance(conf config.Config) Oss {
	switch conf.ServerConfig().Oss {
	case TENCENT:
		return newTencentOss(conf.TencentCosConfig())
	}
	panic("无效的对象存储实例")
}
