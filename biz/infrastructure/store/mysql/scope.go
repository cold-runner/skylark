package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// lark部分
func LarkById(db *MysqlIns, uid string) gen.Condition {
	check(db)
	u, _ := uuid.Parse(uid)
	return db.Query.Lark.ID.Eq(u)
}

func LarkByStuNum(db *MysqlIns, stuNum string) gen.Condition {
	check(db)
	return db.Query.Lark.StuNum.Eq(stuNum)
}

func LarkByPhone(db *MysqlIns, phone string) gen.Condition {
	check(db)
	return db.Query.Lark.Phone.Eq(phone)
}

func LarkByQqUnionId(db *MysqlIns, qqUnionId string) gen.Condition {
	check(db)
	return db.Query.Lark.QqUnionID.Eq(qqUnionId)
}

func LarkQqUnionIdIsNull(db *MysqlIns) gen.Condition {
	check(db)
	return db.Query.Lark.QqUnionID.IsNotNull()
}

// post部分
func PostByUserId(db *MysqlIns, userId uuid.UUID) gen.Condition {
	check(db)
	return db.Query.Post.UserID.Eq(userId.String())
}

// Plate部分
func PlateByName(db *MysqlIns, plateName string) gen.Condition {
	check(db)
	return db.Query.Plate.Name.Eq(plateName)
}

// Category部分
func CategoryByName(db *MysqlIns, cateName string) gen.Condition {
	check(db)
	return db.Query.Categorie.Name.Eq(cateName)
}

// 公共部分
func Paginate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func check(db *MysqlIns) {
	if db == nil || db.Query == nil {
		panic("数据库实例为空！")
	}
}
