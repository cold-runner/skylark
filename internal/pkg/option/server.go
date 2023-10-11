package option

import (
	"time"
)

type Server struct {
	Mode           string        `mapstructure:"mode"`
	Middlewares    []string      `mapstructure:"middlewares"`
	Prefix         string        `mapstructure:"prefix"`
	ExitWaitTime   time.Duration `mapstructure:"exitWaitTime"`
	TrustedProxies []string      `mapstructure:"trustedProxies"`
	Salt           string        `mapstructure:"salt"`

	Sms           string `mapstructure:"sms"`
	SmsNumber     string `mapstructure:"smsNumber"`
	SmsExpiration string `mapstructure:"smsExpiration"`

	Db    string `mapstructure:"store"`
	Cache string `mapstructure:"cache"`
	Oss   string `mapstructure:"oss"`
	Mq    string `mapstructure:"mq"`

	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`

	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string `mapstructure:"cert-file"`
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string `mapstructure:"private-key-file"`
}
