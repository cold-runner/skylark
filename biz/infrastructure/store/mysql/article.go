package mysql

import (
	"context"

	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"gorm.io/gen"
)

func (m *MysqlIns) GetCategoryList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Categorie, error) {
	return m.Query.Categorie.WithContext(c).Where(conds...).Find()
}
