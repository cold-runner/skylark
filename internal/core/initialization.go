package core

import (
	serviceV1 "github.com/cold-runner/skylark/internal/application/v1"
	"github.com/cold-runner/skylark/internal/iface"
	controllerV1 "github.com/cold-runner/skylark/internal/iface/v1"
	"github.com/cold-runner/skylark/internal/infrastructure/cache"
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/internal/infrastructure/db"
	"github.com/cold-runner/skylark/internal/infrastructure/oss"
	"github.com/cold-runner/skylark/internal/infrastructure/sms"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
)

func initController() iface.UserInterface {
	var (
		cacheIns cache.Cache
		ossIns   oss.Oss
		smsIns   sms.Sms
		storeIns store.Store
	)
	// 根据模式初始化组件
	func() {
		globalConf := config.GetConfig()
		serverConf := globalConf.ServerConfig()

		switch {
		// TODO 设置debug模式下的组件
		case serverConf.Mode == config.DEBUG:
			panic("还尚不支持调试模式")
			config.DConfig = &config.DebugConfig{
				Phone:   "11122233344",
				SmsCode: "123456",
			}

		// TODO 设置test模式下的组件
		case serverConf.Mode == config.TEST:
			panic("还尚不支持测试模式")
			fallthrough

		// 生产模式配置
		case serverConf.Mode == config.PRODUCTION:
			cacheIns = cache.NewInstance(globalConf)
			ossIns = oss.NewInstance(globalConf)
			smsIns = sms.NewInstance(globalConf)
			storeIns = db.NewInstance(globalConf)

		default:
			panic("无效的模式，可选项为debug,test,production")
		}
		return
	}()

	serviceIns := serviceV1.NewFactory().NewInstance(
		cacheIns,
		ossIns,
		smsIns,
		storeIns,
	)
	return controllerV1.NewFactory().NewInstance(serviceIns)
}
