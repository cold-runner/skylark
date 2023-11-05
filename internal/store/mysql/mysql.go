package mysql

import (
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/internal/store/mysql/query"
	"github.com/cold-runner/skylark/pkg/db"
)

type mysql struct {
	*query.Query
}

func NewMySQLIns(conf *config.MySQL) store.Store {
	opt := &db.Options{
		Host:                  conf.Host,
		Username:              conf.Username,
		Password:              conf.Password,
		Database:              conf.Database,
		MaxIdleConnections:    conf.MaxIdleConnections,
		MaxOpenConnections:    conf.MaxOpenConnections,
		MaxConnectionLifeTime: conf.MaxConnectionLifeTime,
		LogLevel:              conf.LogLevel,
		Logger:                nil,
	}
	dbIns, err := db.NewMySQL(opt)
	if err != nil {
		panic(err)
	}
	return &mysql{Query: query.Use(dbIns)}
}
