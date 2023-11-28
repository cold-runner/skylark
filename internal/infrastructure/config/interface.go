package config

type Config interface {
	MySQLConfig() *MySQL
	RedisConfig() *Redis
	JwtConfig() *Jwt
	LogConfig() *Log
	ServerConfig() *Server

	TencentSmsConfig() *TencentSms
	QiniuCloudConfig() *QiniuCloud

	TencentCosConfig() *TencentCos
}
