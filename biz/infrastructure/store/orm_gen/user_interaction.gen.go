// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameUserInteraction = "user_interaction"

// UserInteraction 社交关系表
type UserInteraction struct {
	ID         uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();comment:自然主键" json:"id"`                                          // 自然主键
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                                                   // 创建时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:deleted_at,priority:1;comment:删除时间（软删除）" json:"deleted_at"`                           // 删除时间（软删除）
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                                                            // 更新时间
	UserID     string         `gorm:"column:user_id;type:char(36);not null;index:user_interaction_lark_id_fk,priority:1;comment:subject" json:"user_id"`         // subject
	FollowedID string         `gorm:"column:followed_id;type:char(36);not null;index:user_interaction_lark_id_fk2,priority:1;comment:object" json:"followed_id"` // object
}

// TableName UserInteraction's table name
func (*UserInteraction) TableName() string {
	return TableNameUserInteraction
}