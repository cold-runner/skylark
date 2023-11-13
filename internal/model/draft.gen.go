// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDraft = "draft"

// Draft mapped from table <draft>
type Draft struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;comment:自然主键" json:"id"`                    // 自然主键
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`    // 创建时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间（软删除）" json:"deleted_at"`        // 删除时间（软删除）
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`             // 更新时间
	Title     string         `gorm:"column:title;type:varchar(100);not null;comment:草稿标题" json:"title"`          // 草稿标题
	Content   string         `gorm:"column:content;type:text;comment:草稿内容" json:"content"`                       // 草稿内容
	UserID    int64          `gorm:"column:user_id;type:bigint;not null;comment:作者id（考虑性能，不加外键）" json:"user_id"` // 作者id（考虑性能，不加外键）
}

// TableName Draft's table name
func (*Draft) TableName() string {
	return TableNameDraft
}
