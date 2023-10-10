package store

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
)

// Store 存储层接口
type Store interface {
	Article
	Lark
	Comment
}

// Lark 用户模块接口
type Lark interface {
	// Register 新用户注册
	Register(c context.Context, user *user.Lark) error
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
