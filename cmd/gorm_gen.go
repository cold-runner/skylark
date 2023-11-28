package main

import (
	"fmt"
	"github.com/cold-runner/skylark/internal/infrastructure/config"
	"github.com/cold-runner/skylark/pkg/db"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	config.Init()
	mysqlConfig := config.GetConfig().MySQLConfig()
	options := &db.Options{
		Host:                  mysqlConfig.Host,
		Username:              mysqlConfig.Username,
		Password:              mysqlConfig.Password,
		Database:              mysqlConfig.Database,
		MaxIdleConnections:    mysqlConfig.MaxIdleConnections,
		MaxOpenConnections:    mysqlConfig.MaxOpenConnections,
		MaxConnectionLifeTime: mysqlConfig.MaxConnectionLifeTime,
		LogLevel:              mysqlConfig.LogLevel,
		Logger:                nil,
	}
	// 连接数据库
	mysql, err := db.NewMySQL(options)
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 构造生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "internal/store/mysql/query_gen",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		ModelPkgPath:      "internal/domain/model_gen",
		FieldNullable:     false,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	// 设置目标 db
	g.UseDB(mysql)

	// 兼容protobuf
	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(gorm.ColumnType) string { return "int64" },
		"smallint":  func(gorm.ColumnType) string { return "int64" },
		"mediumint": func(gorm.ColumnType) string { return "int64" },
		"bigint":    func(gorm.ColumnType) string { return "int64" },
		"int":       func(gorm.ColumnType) string { return "int64" },
	})

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable(
		gen.FieldGORMTagReg("^id$", func(tag field.GormTag) field.GormTag {
			return tag.
				Append("default", "uuid_generate_v4()").
				Set("type", "uuid")
		}),
		gen.FieldType("id", "uuid.UUID"),
	)...)

	g.Execute()
}
