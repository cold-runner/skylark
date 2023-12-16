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
	g.ApplyInterface(func(LarkAllDetails) {}, g.GenerateModel("lark"))

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

type LarkAllDetails interface {
	// SELECT lark.*, t.followee_count, t.follower_count, count(p.user_id) "post_count",count(e.user_id) "essay_count"
	// from lark left join (select lark.id,count(ui.user_id) "follower_count", count(ui.followed_id) "followee_count" from lark
	//    left join user_interaction ui on lark.id = ui.followed_id
	//    left join user_interaction ui2 on lark.id = ui2.user_id
	//    group by lark.id, ui.user_id, ui.followed_id) t on t.id = lark.id
	// left join skylark.post p on lark.id = p.user_id
	// left join essay e on  lark.id = e.user_id
	// group by lark.id, p.user_id, e.user_id,t.followee_count, t.follower_count
	GetAllDetail() ([]gen.M, error)
	// SELECT lark.*, t.followee_count, t.follower_count, count(p.user_id) "post_count",count(e.user_id) "essay_count"
	// from lark left join (select lark.id,count(ui.user_id) "follower_count", count(ui.followed_id) "followee_count" from lark
	//    left join user_interaction ui on lark.id = ui.followed_id
	//    left join user_interaction ui2 on lark.id = ui2.user_id
	//    group by lark.id, ui.user_id, ui.followed_id) t on t.id = lark.id
	// left join skylark.post p on lark.id = p.user_id
	// left join essay e on  lark.id = e.user_id
	// where lark.@@column=@value
	// group by lark.id, p.user_id, e.user_id,t.followee_count, t.follower_count
	GetAllDetailByField(column string, value string) (gen.M, error)
}
