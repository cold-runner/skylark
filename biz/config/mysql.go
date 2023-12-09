package config

var MySQLConfig = &MySQL{}

type MySQL struct {
	Host                  string `mapstructure:"host"`
	Port                  string `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	Database              string `mapstructure:"database"`
	MaxIdleConnections    int    `mapstructure:"max-idle-connections"`
	MaxOpenConnections    int    `mapstructure:"max-open-connections"`
	MaxConnectionLifeTime int    `mapstructure:"max-connection-life-time"`
	LogMode               int    `mapstructure:"log-level"`
}
