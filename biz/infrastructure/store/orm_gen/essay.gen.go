// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameEssay = "essay"

// Essay 随笔表
type Essay struct {
	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();comment:自然主键" json:"id"`                // 自然主键
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                         // 创建时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:deleted_at,priority:1;comment:删除时间（软删除）" json:"deleted_at"` // 删除时间（软删除）
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                                  // 更新时间
	Content     string         `gorm:"column:content;type:text;not null;comment:随笔内容" json:"content"`                                   // 随笔内容
	TopicID     string         `gorm:"column:topic_id;type:char(36);index:essay_topic_id_fk,priority:1;comment:所属话题" json:"topic_id"`   // 所属话题
	UserID      string         `gorm:"column:user_id;type:char(36);not null;comment:随笔作者（考虑性能，不加外键）" json:"user_id"`                    // 随笔作者（考虑性能，不加外键）
	Temperature int64          `gorm:"column:temperature;type:bigint unsigned;not null;comment:随笔热度（排序时用）" json:"temperature"`          // 随笔热度（排序时用）
	Click       int64          `gorm:"column:click;type:bigint unsigned;not null;comment:点击量（计算热度时用）" json:"click"`                     // 点击量（计算热度时用）
}

// TableName Essay's table name
func (*Essay) TableName() string {
	return TableNameEssay
}
