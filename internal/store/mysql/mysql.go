package mysql

import (
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/store"
	"github.com/cold-runner/skylark/pkg/db"
	"gorm.io/gorm"
)

type mysql struct {
	db *gorm.DB
}

func NewMysqlIns(o *config.MySQL) store.Store {
	options := &db.Options{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
		Logger:                nil,
	}
	dbIns, err := db.NewMySQL(options)
	if err != nil {
		panic(err)
	}
	err = dbIns.AutoMigrate(
		&user.Lark{}, // 用户表
		//&post.StoredLearningPost{}, // 文章表
	)
	if err != nil {
		panic(err)
	}
	return &mysql{db: dbIns}
}
