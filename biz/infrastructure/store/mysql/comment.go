package mysql

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
)

func (m *mysqlIns) CreateComment(c context.Context, comment *orm_gen.Comment) error {
	return m.Query.Comment.WithContext(c).Create(comment)
}
