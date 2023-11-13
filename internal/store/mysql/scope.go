package mysql

import (
	"github.com/cold-runner/skylark/internal/gorm_gen"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// lark部分
func LarkByStuNum(stuNum string) gen.Condition {
	return gorm_gen.Lark.StuNum.Eq(stuNum)
}

func LarkByPhone(phone string) gen.Condition {
	return gorm_gen.Lark.Phone.Eq(phone)
}

func LarkByQqUnionId(qqUnionId string) gen.Condition {
	return gorm_gen.Lark.QqUnionID.Eq(qqUnionId)
}

func LarkQqUnionIdIsNull() gen.Condition {
	return gorm_gen.Lark.QqUnionID.IsNotNull()
}

// post部分
func DraftByUserId(userId int64) gen.Condition {
	return gorm_gen.Draft.UserID.Eq(userId)
}

// 公共部分
func Paginate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
