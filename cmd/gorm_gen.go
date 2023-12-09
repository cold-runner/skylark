package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		"root",
		"my-secret",
		"127.0.0.1:3306",
		"skylark",
		true,
		"Local",
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish store connection: %w", err))
	}
	// 构造生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "biz/infrastructure/store/mysql/query",
		ModelPkgPath:      "biz/infrastructure/store/orm_gen",
		Mode:              gen.WithQueryInterface | gen.WithoutContext, // gen.DefaultQuery
		FieldNullable:     false,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	// 设置目标 store
	g.UseDB(db)

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
