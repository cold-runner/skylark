package mysql

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"gorm.io/gorm"
)

func (m *mysql) CreateLark(c context.Context, lark *user.Lark) error {
	if err := m.db.WithContext(c).Create(lark).Error; err != nil {
		return err
	}
	return nil
}

// GetLarkInfo 利用高阶函数
func (m *mysql) GetLarkInfo(c context.Context, scopes ...func(db *gorm.DB) *gorm.DB) (*user.Lark, error) {
	storedLark := &user.Lark{}
	if err := m.db.WithContext(c).Scopes(scopes...).Find(storedLark).Error; err != nil {
		return nil, err
	}
	return storedLark, nil
}

// UpdateLark 利用高阶函数
func (m *mysql) UpdateLark(c context.Context, values interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if err := m.db.WithContext(c).Model(&user.Lark{}).Scopes(scopes...).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
