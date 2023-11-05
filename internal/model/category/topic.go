package category

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name string // 话题名称

}
