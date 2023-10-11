package config

import "github.com/marmotedu/iam/pkg/log"

type Config interface {
	MySQLConfig() *MySQL
	RedisConfig() *Redis
	JwtConfig() *Jwt
	LogConfig() *log.Options
	ServerConfig() *Server

	TencentSmsConfig() *TencentSms
	QiniuCloudConfig() *QiniuCloud

	TencentCosConfig() *TencentCos
}
