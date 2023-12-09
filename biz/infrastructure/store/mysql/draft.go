package mysql

import (
	"context"

	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func (m *MysqlIns) CreateDraft(c context.Context, draft *orm_gen.Draft) error {
	return m.Query.Draft.WithContext(c).Create(draft)
}

func (m *MysqlIns) GetDraft(c context.Context, conds ...gen.Condition) (*orm_gen.Draft, error) {
	return m.Query.Draft.WithContext(c).Where(conds...).First()
}

func (m *MysqlIns) GetDraftList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Draft, error) {
	return m.Query.Draft.WithContext(c).Where(conds...).Find()
}

func (m *MysqlIns) UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Select(selectScopes...).Where(whereScopes...).Updates(draft)
	return err
}

func (m *MysqlIns) UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Omit(omitScopes...).Where(whereScopes...).Updates(draft)
	return err
}

// ExistDraft 查询是否存在目标用户
func (m *MysqlIns) ExistDraft(c context.Context, conds ...gen.Condition) (bool, error) {
	count, err := m.Query.Draft.WithContext(c).Where(conds...).Count()
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}
