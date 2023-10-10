package comment

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	ParentId    uint
	Content     string
	UserId      uint
	ReplyUserId uint
	PostId      uint
}
