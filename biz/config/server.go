package config

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Mode           string        `mapstructure:"mode"`
	Middlewares    []string      `mapstructure:"middlewares"`
	Prefix         string        `mapstructure:"prefix"`
	ExitWaitTime   time.Duration `mapstructure:"exitWaitTime"`
	TrustedProxies []string      `mapstructure:"trustedProxies"`

	Sms           string `mapstructure:"sms"`
	SmsNumber     int    `mapstructure:"smsNumber"`
	SmsExpiration int    `mapstructure:"smsExpiration"`

	Db           string `mapstructure:"db"`
	Cache        string `mapstructure:"cache"`
	Oss          string `mapstructure:"oss"`
	Mq           string `mapstructure:"mq"`
	SearchEngine string `mapstructure:"searchEngine"`

	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`

	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string `mapstructure:"cert-file"`
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string `mapstructure:"private-key-file"`

	Tls *tls.Config
}

func initTls() {

}
