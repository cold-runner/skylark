package config

import (
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/spf13/viper"
)

var (
	config = new(conf)
)

func init() {
	// 配置文件名为 config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 配置文件路径
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Errorf("read config failed!, err : %v\n", err))
	}

	// TODO 校验配置参数
	if err := viper.Unmarshal(&config); err != nil {
		panic(errors.Errorf("unmarshal config failed!, err : %v\n", err))
	}
}

type conf struct {
	*MySQL  `mapstructure:"mysql"`
	*Redis  `mapstructure:"redis"`
	*Jwt    `mapstructure:"jwt"`
	Log     *log.Options `mapstructure:"log"`
	*Server `mapstructure:"server"`

	*TencentCos `mapstructure:"tencentCos"`
	*QiniuCloud `mapstructure:"qiNiuCloud"`

	*TencentSms `mapstructure:"tencentSms"`
}

// GetConfig 饿汉单例模式初始化全局配置对象
func GetConfig() Config {
	return config
}

func (c *conf) MySQLConfig() *MySQL {
	return c.MySQL
}

func (c *conf) RedisConfig() *Redis {
	return c.Redis
}

func (c *conf) JwtConfig() *Jwt {
	return c.Jwt
}

func (c *conf) LogConfig() *log.Options {
	return c.Log
}

func (c *conf) ServerConfig() *Server {
	return c.Server
}

func (c *conf) TencentSmsConfig() *TencentSms {
	return c.TencentSms
}

func (c *conf) QiniuCloudConfig() *QiniuCloud {
	return c.QiniuCloud
}

func (c *conf) TencentCosConfig() *TencentCos {
	return c.TencentCos
}
