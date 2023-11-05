package store

import (
	"context"
	"gorm.io/gen/field"

	"github.com/cold-runner/skylark/internal/model"
	"gorm.io/gen"
)

// Store 存储层接口
type Store interface {
	Post
	Lark
	Comment
	Exist(c context.Context, conds ...gen.Condition) (bool, error)
}

// 数据库依赖大概率不会发生改变，所以这里定义依赖为针对于mysql的可变参数

// Lark 用户模块接口
type Lark interface {
	CreateLark(c context.Context, user *model.Lark) error
	GetLark(c context.Context, conds ...gen.Condition) (*model.Lark, error)
	GetLarkList(c context.Context, conds ...gen.Condition) ([]*model.Lark, error)
	UpdateLarkSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, lark *model.Lark) error
	UpdateLarkOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, lark *model.Lark) error
}

// Comment 评论模块接口
type Comment interface {
	// PostComment 发布评论
	//PostComment(c context.Context)
}

// Post 文章模块接口
type Post interface {
	CreateDraft(c context.Context, draft *model.Draft) error
	GetDraft(c context.Context, conds ...gen.Condition) (*model.Draft, error)
	GetDraftList(c context.Context, conds ...gen.Condition) ([]*model.Draft, error)
	UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, draft *model.Draft) error
	UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *model.Draft) error
}
