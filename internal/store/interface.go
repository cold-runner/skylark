package store

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"gorm.io/gorm"
)

// Store 存储层接口
type Store interface {
	Article
	Lark
	Comment
}

// Lark 用户模块接口
type Lark interface {
	CreateLark(c context.Context, user *user.Lark) error
	GetLarkInfo(c context.Context, scopes ...func(db *gorm.DB) *gorm.DB) (*user.Lark, error)
	UpdateLark(c context.Context, values interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
}

// Comment 评论模块接口
type Comment interface {
	// PostComment 发布评论
	//PostComment(c context.Context)
}

// Article 文章模块接口
type Article interface {
	// PostArticle 发表一篇文章
	//PostArticle(c context.Context)
}
