package post

import "gorm.io/gorm"

type StoredDraft struct {
	gorm.Model
	Title   string
	Content string
	UserId  uint
}
