package mysql

import (
	"context"

	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func (m *mysqlIns) CreateLark(c context.Context, user *orm_gen.Lark) error {
	return m.Query.Lark.WithContext(c).Create(user)
}

func (m *mysqlIns) GetLark(c context.Context, conds ...gen.Condition) (*orm_gen.Lark, error) {
	return m.Query.Lark.WithContext(c).Where(conds...).First()
}

func (m *mysqlIns) GetLarkList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Lark, error) {
	return m.Query.Lark.WithContext(c).Where(conds...).Find()
}

func (m *mysqlIns) UpdateLarkSelect(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *orm_gen.Lark) error {
	_, err := m.Query.Lark.WithContext(c).Select(selectScopes...).Where(whereScopes...).Updates(lark)
	return err
}

func (m *mysqlIns) UpdateLarkOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, lark *orm_gen.Lark) error {
	_, err := m.Query.Lark.WithContext(c).Omit(omitScopes...).Where(whereScopes...).Updates(lark)
	return err
}

// ExistUser 查询是否存在目标用户
func (m *mysqlIns) ExistUser(c context.Context, conds ...gen.Condition) (bool, error) {
	count, err := m.Query.Lark.WithContext(c).Where(conds...).Count()
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}
