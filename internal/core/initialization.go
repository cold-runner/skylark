package core

import (
	"github.com/cold-runner/skylark/internal/controller"
	controllerV1 "github.com/cold-runner/skylark/internal/controller/v1"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/pkg/db"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	serviceV1 "github.com/cold-runner/skylark/internal/service/v1"
	"github.com/cold-runner/skylark/internal/store"
)

func initController() controller.Interface {
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
			config.DConfig = &config.DebugConfig{
				Phone:   "11122233344",
				SmsCode: "123456",
			}

		// TODO 设置test模式下的组件
		case serverConf.Mode == config.TEST:
			fallthrough

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
	return controllerV1.NewFactory().NewInstance(nil, serviceIns)
}
