// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID         int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:idx_tags_deleted_at,priority:1" json:"deleted_at"`
	Name       string         `gorm:"column:name;type:longtext" json:"name"`
	CategoryID int64          `gorm:"column:category_id;type:bigint unsigned" json:"category_id"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
