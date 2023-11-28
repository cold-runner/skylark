package mysql

import (
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
	"github.com/cold-runner/skylark/internal/infrastructure/store/mysql/query_gen"
	"github.com/cold-runner/skylark/pkg/db"
)

type mysql struct {
	*query_gen.Query
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
	query_gen.SetDefault(dbIns)
	return &mysql{Query: query_gen.Use(dbIns)}
}
