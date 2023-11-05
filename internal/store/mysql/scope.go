package mysql

import (
	"github.com/cold-runner/skylark/internal/store/mysql/query"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// lark部分
func LarkByStuNum(stuNum string) gen.Condition {
	return query.Lark.StuNum.Eq(stuNum)
}

func LarkByPhone(phone string) gen.Condition {
	return query.Lark.Phone.Eq(phone)
}

func LarkByQqUnionId(qqUnionId string) gen.Condition {
	return query.Lark.QqUnionID.Eq(qqUnionId)
}

func LarkQqUnionIdIsNull() gen.Condition {
	return query.Lark.QqUnionID.IsNotNull()
}

// post部分
func DraftByUserId(userId int64) gen.Condition {
	return query.Draft.UserID.Eq(userId)
}

// 公共部分
func Paginate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
