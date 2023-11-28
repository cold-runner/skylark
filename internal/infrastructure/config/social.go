package config

type QqSocial struct {
	AppId        string `mapstructure:"appId"`
	ClientSecret string `mapstructure:"clientSecret"`
	redirectUrl  string `mapstructure:"redirectUrl"`
}
