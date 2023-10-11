package mysql

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
)

func (m *mysql) Create(c context.Context, lark *user.Lark) error {
	if err := m.db.WithContext(c).Create(lark).Error; err != nil {
		return err
	}
	return nil
}
