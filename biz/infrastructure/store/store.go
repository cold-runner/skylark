package store

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type Type string

func (s Type) String() string {
	return string(s)
}

const (
	MYSQL Type = "mysql"
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

func GetIns() Store {
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
	Lark
	LarkScope

	Post
	PostScope

	Comment
	CommentScope
}

// Lark 用户模块接口
type Lark interface {
	CreateLark(c context.Context, user *orm_gen.Lark) error
	GetLark(c context.Context, conds ...gen.Condition) (*orm_gen.Lark, error)
	GetLarkAllDetail(column, value string) (map[string]interface{}, error)
	IsFollowed(c context.Context, userId, tarId string) (bool, error)
	GetLarkList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Lark, error)
	LarkListAllDetail() ([]map[string]interface{}, error)
	UpdateLarkSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, lark *orm_gen.Lark) error
	UpdateLarkOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, lark *orm_gen.Lark) error
}
type LarkScope interface {
	LarkById(uid string) gen.Condition
	LarkByStuNum(stuNum string) gen.Condition
	LarkByPhone(phone string) gen.Condition
	LarkByQqUnionId(qqUnionId string) gen.Condition
}

// Comment 评论模块接口
type Comment interface {
	// CreateComment 发布评论
	CreateComment(c context.Context, comment *orm_gen.Comment) error
	UpdateComment(c context.Context, commentId, content string) error
	DeleteComment(c context.Context, commentId string) error
	QueryComment(c context.Context, postId string, idx, limit int) ([]*orm_gen.Comment, error)
	QueryReplyComment(c context.Context, postId, parentId string, idx, limit int) ([]*orm_gen.Comment, error)
}

type CommentScope interface {
}

// Post 文章模块接口
type Post interface {
	CreateDraft(c context.Context, draft *orm_gen.Draft) error
	GetDraft(c context.Context, conds ...gen.Condition) (*orm_gen.Draft, error)
	GetDraftList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Draft, error)
	UpdateDraftSelect(c context.Context, selectScopes []field.Expr, whereScope []gen.Condition, draft *orm_gen.Draft) error
	UpdateDraftOmit(c context.Context, omitScopes []field.Expr, whereScopes []gen.Condition, draft *orm_gen.Draft) error
	GetCategoryList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Categorie, error)
	GetPost(c context.Context, conds ...gen.Condition) (*orm_gen.Post, error)
	GetPostList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Post, error)
	GetTagList(c context.Context, conds ...gen.Condition) ([]*orm_gen.Tag, error)
	GetPostTagList(c context.Context, conds ...gen.Condition) ([]*orm_gen.PostTag, error)
	GetPostListAllDetail(c context.Context) ([]byte, error)
}
type PostScope interface {
	PostById(uid string) gen.Condition
	PostByUserId(userId string) gen.Condition
}
