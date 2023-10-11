package store

import (
	"github.com/cold-runner/skylark/internal/pkg/option"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql"
)

const (
	MYSQL = "mysql"
)

// 简单工厂模式

func NewStoreIns(conf *option.Conf) store.Store {
	switch conf.Server.Db {
	case MYSQL:
		return mysql.NewMysqlIns(conf.MySQL)
	}
	panic("无效的store依赖")
}
