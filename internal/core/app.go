package core

import (
	"crypto/tls"
	"github.com/cloudwego/hertz/pkg/app/server"
	hzConfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cold-runner/skylark/internal/controller"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/pkg/errors"
)

var App = new(Application)

type Application struct {
	router        *server.Hertz
	controllerIns controller.Interface
}

func InitApp() {
	App.controllerIns = initController()

	globalConfig := config.GetConfig()
	log.Init(globalConfig.LogConfig())

	var options []hzConfig.Option
	serverConf := globalConfig.ServerConfig()
	// 设置tls，配置文件留空则不使用tls
	if serverTlsConfig, err := newServerTlsConfig(serverConf.CertFile, serverConf.KeyFile); err == nil {
		options = append(options, serverTlsConfig)
	} else {
		log.Warnf("未开启TLS")
	}

	// 设置监听地址、端口
	options = append(options, server.WithHostPorts(serverConf.Host+":"+serverConf.Port))
	// 设置退出时间
	options = append(options, server.WithExitWaitTime(serverConf.ExitWaitTime))

	App.router = server.New(options...)
}

func newServerTlsConfig(certFilePath, keyFilePath string) (hzConfig.Option, error) {
	if certFilePath == "" || keyFilePath == "" {
		return hzConfig.Option{}, errors.New("证书或密钥文件为空，不使用tls配置")
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
	a.router.Use(gzip.Gzip(gzip.DefaultCompression), accesslog.New())
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
