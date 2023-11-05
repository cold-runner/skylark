package mysql

import (
	"context"
	"gorm.io/gen/field"

	"github.com/cold-runner/skylark/internal/model"
	"gorm.io/gen"
)

func (m *mysql) CreateLark(c context.Context, lark *model.Lark) error {
	if err := m.Lark.WithContext(c).Create(lark); err != nil {
		return err
	}
	return nil
}

// GetLark 利用高阶函数
func (m *mysql) GetLark(c context.Context, conds ...gen.Condition) (*model.Lark, error) {
	lark, err := m.Lark.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return lark, nil
}

// GetLarkList 利用高阶函数
func (m *mysql) GetLarkList(c context.Context, conds ...gen.Condition) ([]*model.Lark, error) {
	larks, err := m.Lark.WithContext(c).Where(conds...).Find()
	if err != nil {
		return nil, err
	}
	return larks, nil
}

// UpdateLarkSelect 利用高阶函数
func (m *mysql) UpdateLarkSelect(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error {
	if _, err := m.Query.Lark.WithContext(c).Select(selectScopes...).Where(whereScopes...).Updates(lark); err != nil {
		return err
	}
	return nil
}

// UpdateLarkOmit 利用高阶函数
func (m *mysql) UpdateLarkOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error {
	if _, err := m.Query.Lark.WithContext(c).Omit(omitScopes...).Where(whereScopes...).Updates(lark); err != nil {
		return err
	}
	return nil
}

// Exist 查询是否存在目标用户
func (m *mysql) Exist(c context.Context, conds ...gen.Condition) (bool, error) {
	count, err := m.Query.Lark.WithContext(c).Where(conds...).Count()
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}
