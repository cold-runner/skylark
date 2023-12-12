package mysql

import (
	"github.com/cold-runner/skylark/biz/infrastructure/store/mysql/query"
	"time"

	"github.com/cold-runner/skylark/biz/config"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/pkg/db"
)

type mysqlIns struct {
	*query.Query
}

func NewFactory() store.Factory {
	return new(mysqlIns)
}

func (m *mysqlIns) NewInstance() store.Store {
	conf := config.GetConfig().GetMySQL()
	cfg := &db.Options{
		Host:                  conf.Host + ":" + conf.Port,
		Username:              conf.User,
		Password:              conf.Password,
		Database:              conf.Database,
		MaxIdleConnections:    conf.MaxIdleConnections,
		MaxOpenConnections:    conf.MaxOpenConnections,
		MaxConnectionLifeTime: time.Duration(conf.MaxConnectionLifeTime),
		LogLevel:              conf.LogMode,
		Logger:                nil,
	}
	dbIns, err := db.NewMySQL(cfg)
	if err != nil {
		panic(err)
	}

	return &mysqlIns{query.Use(dbIns)}
}
