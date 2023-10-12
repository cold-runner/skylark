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

func (m *mysql) QueryByStuNum(c context.Context, stuNum string) (*user.Lark, error) {
	storedUser := &user.Lark{}
	if err := m.db.WithContext(c).Where("stu_num = ?", stuNum).First(storedUser).Error; err != nil {
		return nil, err
	}
	return storedUser, nil
}

func (m *mysql) QueryByPhone(c context.Context, phone string) (*user.Lark, error) {
	storedUser := &user.Lark{}
	if err := m.db.WithContext(c).Where("phone = ?", phone).First(storedUser).Error; err != nil {
		return nil, err
	}
	return storedUser, nil
}

func (m *mysql) QueryByQqUnionId(c context.Context, qq string) (*user.Lark, error) {
	storedUser := &user.Lark{}
	if err := m.db.WithContext(c).Where("qq_union_id = ?", qq).First(storedUser).Error; err != nil {
		return nil, err
	}
	return storedUser, nil
}
