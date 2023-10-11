package oss

import "github.com/cold-runner/skylark/internal/pkg/option"

// 简单工厂模式

func NewOss(conf *option.Conf) Oss {
	switch conf.Server.Oss {
	case TENCENT:
		return NewTencentOss(conf.TencentCos)
	}
	panic("无效的对象存储实例")
}
