package entity

import "gorm.io/gorm"

type Plate struct {
	gorm.Model
	Name string // 板块名称

}
