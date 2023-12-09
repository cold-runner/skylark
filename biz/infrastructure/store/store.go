package store

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type StoreType string

func (s StoreType) String() string {
	return string(s)
}

const (
	MYSQL StoreType = "mysql"
)

var (
	storeIns Store
)

func SetDefault(s Store) {
	if s == nil {
		panic("store is nil")
	}
	storeIns = s
}

func GetStoreIns() Store {
	if storeIns == nil {
		panic("store is nil")
	}
	return storeIns
}

// Factory 工厂方法模式
type Factory interface {
	NewInstance() Store
}

// Store 存储层接口
type Store interface {
	Post
	Lark
	Comment
}

// 数据库依赖大概率不会发生改变，所以这里定义依赖为针对于mysql的可变参数

// Lark 用户模块接口
type Lark interface {
	CreateLark(c context.Context, user *orm_gen.Lark) error
	GetLark(c context.Context, conds ...gen.Condition) (*orm_gen.Lark, error)
	GetLarkList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Lark, error)
	UpdateLarkSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, lark *orm_gen.Lark) error
	UpdateLarkOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, lark *orm_gen.Lark) error
}

// Comment 评论模块接口
type Comment interface {
	// PostComment 发布评论
	//PostComment(c context.Context)
}

// Post 文章模块接口
type Post interface {
	CreateDraft(c context.Context, draft *orm_gen.Draft) error
	GetDraft(c context.Context, conds ...gen.Condition) (*orm_gen.Draft, error)
	GetDraftList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Draft, error)
	UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, draft *orm_gen.Draft) error
	UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error
}
