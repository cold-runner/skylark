// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameTag = "tag"

// Tag 标签表
type Tag struct {
	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();comment:自然主键" json:"id"`                            // 自然主键
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                                     // 创建时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:deleted_at,priority:1;comment:删除时间（软删除）" json:"deleted_at"`             // 删除时间（软删除）
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                                              // 更新时间
	Name        string         `gorm:"column:name;type:varchar(20);not null;uniqueIndex:tag_pk2,priority:1;comment:标签名" json:"name"`                // 标签名
	CategorieID string         `gorm:"column:categorie_id;type:char(36);index:tag_categorie_id_fk,priority:1;comment:该标签隶属的归档" json:"categorie_id"` // 该标签隶属的归档
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}