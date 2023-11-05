package user

import "gorm.io/gorm"

type UserInteract struct {
	gorm.Model
	UserId       uint
	FollowUserId uint
}
