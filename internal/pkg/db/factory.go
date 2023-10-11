package db

import (
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql"
)

const (
	MYSQL = "mysql"
)

// 简单工厂模式

func NewInstance(conf config.Config) store.Store {
	switch conf.ServerConfig().Db {
	case MYSQL:
		return mysql.NewMysqlIns(conf.MySQLConfig())
	}
	panic("无效的store依赖")
}
