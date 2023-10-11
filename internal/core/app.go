package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cold-runner/skylark/internal/controller"
	controllerV1 "github.com/cold-runner/skylark/internal/controller/v1"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/pkg/store"
	service "github.com/cold-runner/skylark/internal/service/v1"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var App = new(Application)

type Application struct {
	server     *option.Server
	router     *server.Hertz
	controller controller.Controller
}

func DependencyInjection() {
	appConfig := option.NewConfig()
	// TODO 校验配置参数
	smsNumber, err := strconv.Atoi(appConfig.Server.SmsNumber)
	if err != nil {
		panic(err)
	}
	smsExpiration, err := strconv.Atoi(appConfig.Server.SmsExpiration)
	if err != nil {
		panic(err)
	}
	paramMap := map[string]interface{}{
		sms.SMS_NUMBER:     smsNumber,
		sms.SMS_EXPIRATION: smsExpiration,
	}
	ctx := context.WithValue(context.Background(), "paramMap", paramMap)

	// 根据配置文件注入所有依赖
	App.controller = controllerV1.NewControllerV1(
		ctx,
		service.NewServiceV1(
			cache.NewCache(appConfig),
			oss.NewOss(appConfig),
			sms.NewSmsClient(appConfig),
			store.NewStoreIns(appConfig),
		),
	)

	var options []config.Option

	// 设置tls，配置文件留空则不使用tls
	if serverTlsConfig, err := newServerTlsConfig(appConfig.CertFile, appConfig.KeyFile); err == nil {
		options = append(options, serverTlsConfig)
	} else {
		fmt.Println("未启用tls")
	}

	// 设置监听地址、端口
	options = append(options, server.WithHostPorts(appConfig.Server.Host+":"+appConfig.Server.Port))
	// 设置退出时间
	options = append(options, server.WithExitWaitTime(appConfig.ExitWaitTime))

	App.router = server.New(options...)
}

func newServerTlsConfig(certFilePath, keyFilePath string) (config.Option, error) {
	if certFilePath == "" || keyFilePath == "" {
		return config.Option{}, errors.New("证书或密钥文件为空，不使用tls配置")
	}
	cfg := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		log.Panicf("配置tls证书失败， err: %v", err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	return server.WithTLS(cfg), nil
}

func (a *Application) InstallRouter() *Application {
	// 注册中间件
	a.router.Use(accesslog.New())
	// 公共路由
	a.publicRouter()
	// 文章路由
	a.articleRouter()
	// 用户路由
	a.larkRouter()
	// 评论路由
	a.commentRouter()

	return a
}

func (a *Application) Run() {
	a.router.Spin()
}
