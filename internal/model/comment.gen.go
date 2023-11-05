// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID          int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:idx_comments_deleted_at,priority:1" json:"deleted_at"`
	ParentID    int64          `gorm:"column:parent_id;type:bigint unsigned" json:"parent_id"`
	Content     string         `gorm:"column:content;type:longtext" json:"content"`
	UserID      int64          `gorm:"column:user_id;type:bigint unsigned" json:"user_id"`
	ReplyUserID int64          `gorm:"column:reply_user_id;type:bigint unsigned" json:"reply_user_id"`
	PostID      int64          `gorm:"column:post_id;type:bigint unsigned" json:"post_id"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
