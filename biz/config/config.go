package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	DEBUG      = MODE("debug")
	TEST       = MODE("test")
	PRODUCTION = MODE("production")
)

var (
	c Conf
)

func GetConfig() Conf {
	if c == nil {
		panic(errors.New("config not init"))
	}
	return c
}

type Conf interface {
	GetMode() MODE
	GetMySQL() *MySQL
	GetRedis() *Redis
	GetJwt() *Jwt
	GetLog() *Log
	GetServer() *Server
	GetTencentCos() *TencentCos
	GetTencentSms() *TencentSms
	//GetQiniu() *Qiniu
}

type MODE string

func (m MODE) String() string {
	return string(m)
}

type config struct {
	*MySQL  `mapstructure:"mysql"`
	*Redis  `mapstructure:"redis"`
	*Jwt    `mapstructure:"jwt"`
	*Log    `mapstructure:"log"`
	*Server `mapstructure:"server"`

	*TencentCos `mapstructure:"tencentCos"`
	*QiniuCloud `mapstructure:"qiNiuCloud"`

	*TencentSms `mapstructure:"tencentSms"`

	*QqSocial `mapstructure:"qqSocial"`
}

func (c *config) GetTencentSms() *TencentSms {
	return c.TencentSms
}

func (c *config) GetMode() MODE {
	return MODE(c.Mode)
}

func (c *config) GetMySQL() *MySQL {
	return c.MySQL
}

func (c *config) GetRedis() *Redis {
	return c.Redis
}

func (c *config) GetJwt() *Jwt {
	return c.Jwt
}

func (c *config) GetLog() *Log {
	return c.Log
}

func (c *config) GetServer() *Server {
	return c.Server
}

func (c *config) GetTencentCos() *TencentCos {
	return c.TencentCos
}

func Init() {
	// 配置文件名为 config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 配置文件路径
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Errorf("read config failed!, err : %v\n", err))
	}
	c = new(config)
	if err := viper.Unmarshal(c); err != nil {
		panic(errors.Errorf("unmarshal config failed!, err : %v\n", err))
	}
}
