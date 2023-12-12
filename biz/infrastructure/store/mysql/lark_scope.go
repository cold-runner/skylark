package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gen"
)

func (m *mysqlIns) LarkById(id string) gen.Condition {
	uid, _ := uuid.Parse(id)
	return m.Query.Lark.ID.Eq(uid)
}

func (m *mysqlIns) LarkByStuNum(stuNum string) gen.Condition {
	return m.Query.Lark.StuNum.Eq(stuNum)
}

func (m *mysqlIns) LarkByPhone(phone string) gen.Condition {
	return m.Query.Lark.Phone.Eq(phone)
}

func (m *mysqlIns) LarkByQqUnionId(qqUnionId string) gen.Condition {
	return m.Query.Lark.QqUnionID.Eq(qqUnionId)
}

func (m *mysqlIns) PostById(id string) gen.Condition {
	uid, _ := uuid.Parse(id)
	return m.Query.Post.ID.Eq(uid)
}

func (m *mysqlIns) PostByUserId(userId string) gen.Condition {
	return m.Query.Post.UserID.Eq(userId)
}
