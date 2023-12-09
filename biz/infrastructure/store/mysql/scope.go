package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// lark部分
func LarkByStuNum(db *mysqlIns, stuNum string) gen.Condition {
	check(db)
	return db.Query.Lark.StuNum.Eq(stuNum)
}

func LarkByPhone(db *mysqlIns, phone string) gen.Condition {
	check(db)
	return db.Query.Lark.Phone.Eq(phone)
}

func LarkByQqUnionId(db *mysqlIns, qqUnionId string) gen.Condition {
	check(db)
	return db.Query.Lark.QqUnionID.Eq(qqUnionId)
}

func LarkQqUnionIdIsNull(db *mysqlIns) gen.Condition {
	check(db)
	return db.Query.Lark.QqUnionID.IsNotNull()
}

// post部分
func PostByUserId(db *mysqlIns, userId uuid.UUID) gen.Condition {
	check(db)
	return db.Query.Post.UserID.Eq(userId.String())
}

// Plate部分
func PlateByName(db *mysqlIns, plateName string) gen.Condition {
	check(db)
	return db.Query.Plate.Name.Eq(plateName)
}

// Category部分
func CategoryByName(db *mysqlIns, cateName string) gen.Condition {
	check(db)
	return db.Query.Categorie.Name.Eq(cateName)
}

// 公共部分
func Paginate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func check(db *mysqlIns) {
	if db == nil || db.Query == nil {
		panic("数据库实例为空！")
	}
}
