package mysql

import (
	"context"
	"gorm.io/gen/field"

	"github.com/cold-runner/skylark/internal/domain/model_gen"
	"gorm.io/gen"
)

func (m *mysql) CreateDraft(c context.Context, draft *model_gen.Draft) error {
	if err := m.Query.Draft.WithContext(c).Create(draft); err != nil {
		return err
	}
	return nil
}

func (m *mysql) GetDraft(c context.Context, conds ...gen.Condition) (*model_gen.Draft, error) {
	draft, err := m.Query.Draft.WithContext(c).Where(conds...).First()
	if err != nil {
		return nil, err
	}
	return draft, nil
}

func (m *mysql) GetDraftList(c context.Context, conds ...gen.Condition) ([]*model_gen.Draft, error) {
	draftList, err := m.Query.WithContext(c).Draft.Where(conds...).Find()
	if err != nil {
		return nil, err
	}
	return draftList, nil
}

func (m *mysql) UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScopes []gen.Condition, draft *model_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Select(selectScopes...).Where(whereScopes...).Updates(draft)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysql) UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *model_gen.Draft) error {
	_, err := m.Query.Draft.WithContext(c).Omit(omitScopes...).Where(whereScopes...).Updates(draft)
	if err != nil {
		return err
	}
	return nil
}
