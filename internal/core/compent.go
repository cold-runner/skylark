package core

import (
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/db"
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql"
)

func newCache(conf *option.Conf) cache.Cache {
	switch conf.Server.Cache {
	case cache.REDIS:
		return cache.NewRedis(conf.Redis)
	}
	panic("无效的缓存实例")
}

func newOss(conf *option.Conf) oss.Oss {
	switch conf.Server.Oss {
	case oss.TENCENT:
		return oss.NewTencentOss(conf.TencentCos)
	}
	panic("无效的对象存储实例")
}

func newSmsClient(conf *option.Conf) sms.Sms {
	switch conf.Server.Sms {
	case sms.TENCENT:
		return sms.NewTencentSms(conf.TencentSms)
	}
	panic("无效的sms服务商")
}

func newStore(conf *option.Conf) store.Store {
	switch conf.Server.Db {
	case db.MYSQL:
		return mysql.NewMysqlIns(conf.MySQL)
	}
	panic("无效的store依赖")
}
