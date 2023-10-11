package option

import (
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/spf13/viper"
)

type Conf struct {
	*MySQL      `mapstructure:"mysql"`
	*Redis      `mapstructure:"redis"`
	*Jwt        `mapstructure:"jwt"`
	Log         *log.Options `mapstructure:"log"`
	*Server     `mapstructure:"server"`
	*QiniuCloud `mapstructure:"qiNiuCloud"`
	*TencentSms `mapstructure:"tencentSms"`
	*TencentCos `mapstructure:"tencentCos"`
}

func NewConfig() *Conf {
	// 配置文件名为 config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 配置文件路径
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Errorf("read config failed!, err : %v\n", err))
	}
	c := &Conf{}
	if err := viper.Unmarshal(c); err != nil {
		panic(errors.Errorf("unmarshal config failed!, err : %v\n", err))
	}
	return c
}
