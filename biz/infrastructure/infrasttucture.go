package infrastructure

import (
	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/handler/user"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/log"
	"github.com/cold-runner/skylark/biz/infrastructure/oss"
	"github.com/cold-runner/skylark/biz/infrastructure/searchEngine"
	"github.com/cold-runner/skylark/biz/infrastructure/sms"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/mysql"
	"github.com/pkg/errors"
)

func Init() {
	c := config.GetConfig()

	// TODO 校验配置参数
	switch c.GetServer().Mode {
	case config.PRODUCTION.String():
		// 初始化仓储层
		func() {
			switch c.GetServer().Db {
			case store.MYSQL.String():
				store.SetDefault(mysql.NewFactory().NewInstance())
			default:
				panic(errors.Errorf("无效的仓储实例类型，支持的类型：mysql。当前传入类型为：%s", c.GetServer().Db))
			}
		}()
		// 初始化认证、授权中间件
		user.InitAuthenticatorAndAuthorization()
		// 初始化日志
		log.Init(c)
		// 初始化缓存
		cache.Init(c)
		// 初始化对象存储
		oss.Init(c)
		// 初始化sms客户端
		sms.Init(c)
		// 初始化全文搜索引擎
		searchEngine.Init(c)
	case config.DEBUG.String():

	case config.TEST.String():
	}
}
