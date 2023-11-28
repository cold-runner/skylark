package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	BackGroundUrl string // 背景图URL
	Name          string // 归档名
	Rank          int    // 排名

	PlateId uint
}
