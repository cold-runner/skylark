package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string // 标签名

	CategoryId uint // 归档
}
