package core

import (
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	hzConfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/iface"
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/internal/infrastructure/log"
	"github.com/pkg/errors"
)

var App = new(Application)

type Application struct {
	router        *server.Hertz
	userInterface iface.UserInterface
}

func InitApp() {
	config.Init()
	App.userInterface = initController()

	globalConfig := config.GetConfig()
	log.Init()
	var options []hzConfig.Option
	serverConf := globalConfig.ServerConfig()
	// 设置tls，配置文件留空则不使用tls
	if serverTlsConfig, err := newServerTlsConfig(serverConf.CertFile, serverConf.KeyFile); err == nil {
		options = append(options, serverTlsConfig)
	} else {
		hlog.Debug("未开启TLS配置")
	}

	// 注册自定义校验规则
	options = append(options, server.WithValidateConfig(iface.Validator()))
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
		fmt.Printf("配置tls证书失败， err: %v", err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	return server.WithTLS(cfg), nil
}

func (a *Application) Run() {
	a.router.Spin()
}
