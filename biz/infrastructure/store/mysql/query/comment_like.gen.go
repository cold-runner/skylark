// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
)

func newCommentLike(db *gorm.DB, opts ...gen.DOOption) commentLike {
	_commentLike := commentLike{}

	_commentLike.commentLikeDo.UseDB(db, opts...)
	_commentLike.commentLikeDo.UseModel(&orm_gen.CommentLike{})

	tableName := _commentLike.commentLikeDo.TableName()
	_commentLike.ALL = field.NewAsterisk(tableName)
	_commentLike.ID = field.NewField(tableName, "id")
	_commentLike.CommentID = field.NewString(tableName, "comment_id")
	_commentLike.UserID = field.NewString(tableName, "user_id")

	_commentLike.fillFieldMap()

	return _commentLike
}

// commentLike 评论点赞表
type commentLike struct {
	commentLikeDo

	ALL       field.Asterisk
	ID        field.Field  // 自然主键
	CommentID field.String // 评论id
	UserID    field.String // 评论者id

	fieldMap map[string]field.Expr
}

func (c commentLike) Table(newTableName string) *commentLike {
	c.commentLikeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentLike) As(alias string) *commentLike {
	c.commentLikeDo.DO = *(c.commentLikeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentLike) updateTableName(table string) *commentLike {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CommentID = field.NewString(table, "comment_id")
	c.UserID = field.NewString(table, "user_id")

	c.fillFieldMap()

	return c
}

func (c *commentLike) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentLike) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["comment_id"] = c.CommentID
	c.fieldMap["user_id"] = c.UserID
}

func (c commentLike) clone(db *gorm.DB) commentLike {
	c.commentLikeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commentLike) replaceDB(db *gorm.DB) commentLike {
	c.commentLikeDo.ReplaceDB(db)
	return c
}

type commentLikeDo struct{ gen.DO }

type ICommentLikeDo interface {
	gen.SubQuery
	Debug() ICommentLikeDo
	WithContext(ctx context.Context) ICommentLikeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentLikeDo
	WriteDB() ICommentLikeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentLikeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentLikeDo
	Not(conds ...gen.Condition) ICommentLikeDo
	Or(conds ...gen.Condition) ICommentLikeDo
	Select(conds ...field.Expr) ICommentLikeDo
	Where(conds ...gen.Condition) ICommentLikeDo
	Order(conds ...field.Expr) ICommentLikeDo
	Distinct(cols ...field.Expr) ICommentLikeDo
	Omit(cols ...field.Expr) ICommentLikeDo
	Join(table schema.Tabler, on ...field.Expr) ICommentLikeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentLikeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentLikeDo
	Group(cols ...field.Expr) ICommentLikeDo
	Having(conds ...gen.Condition) ICommentLikeDo
	Limit(limit int) ICommentLikeDo
	Offset(offset int) ICommentLikeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentLikeDo
	Unscoped() ICommentLikeDo
	Create(values ...*orm_gen.CommentLike) error
	CreateInBatches(values []*orm_gen.CommentLike, batchSize int) error
	Save(values ...*orm_gen.CommentLike) error
	First() (*orm_gen.CommentLike, error)
	Take() (*orm_gen.CommentLike, error)
	Last() (*orm_gen.CommentLike, error)
	Find() ([]*orm_gen.CommentLike, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentLike, err error)
	FindInBatches(result *[]*orm_gen.CommentLike, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.CommentLike) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentLikeDo
	Assign(attrs ...field.AssignExpr) ICommentLikeDo
	Joins(fields ...field.RelationField) ICommentLikeDo
	Preload(fields ...field.RelationField) ICommentLikeDo
	FirstOrInit() (*orm_gen.CommentLike, error)
	FirstOrCreate() (*orm_gen.CommentLike, error)
	FindByPage(offset int, limit int) (result []*orm_gen.CommentLike, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentLikeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentLikeDo) Debug() ICommentLikeDo {
	return c.withDO(c.DO.Debug())
}

func (c commentLikeDo) WithContext(ctx context.Context) ICommentLikeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentLikeDo) ReadDB() ICommentLikeDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentLikeDo) WriteDB() ICommentLikeDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentLikeDo) Session(config *gorm.Session) ICommentLikeDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentLikeDo) Clauses(conds ...clause.Expression) ICommentLikeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentLikeDo) Returning(value interface{}, columns ...string) ICommentLikeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentLikeDo) Not(conds ...gen.Condition) ICommentLikeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentLikeDo) Or(conds ...gen.Condition) ICommentLikeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentLikeDo) Select(conds ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentLikeDo) Where(conds ...gen.Condition) ICommentLikeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentLikeDo) Order(conds ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentLikeDo) Distinct(cols ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentLikeDo) Omit(cols ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentLikeDo) Join(table schema.Tabler, on ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentLikeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentLikeDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentLikeDo) Group(cols ...field.Expr) ICommentLikeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentLikeDo) Having(conds ...gen.Condition) ICommentLikeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentLikeDo) Limit(limit int) ICommentLikeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentLikeDo) Offset(offset int) ICommentLikeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentLikeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentLikeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentLikeDo) Unscoped() ICommentLikeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentLikeDo) Create(values ...*orm_gen.CommentLike) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentLikeDo) CreateInBatches(values []*orm_gen.CommentLike, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentLikeDo) Save(values ...*orm_gen.CommentLike) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentLikeDo) First() (*orm_gen.CommentLike, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentLike), nil
	}
}

func (c commentLikeDo) Take() (*orm_gen.CommentLike, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentLike), nil
	}
}

func (c commentLikeDo) Last() (*orm_gen.CommentLike, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentLike), nil
	}
}

func (c commentLikeDo) Find() ([]*orm_gen.CommentLike, error) {
	result, err := c.DO.Find()
	return result.([]*orm_gen.CommentLike), err
}

func (c commentLikeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentLike, err error) {
	buf := make([]*orm_gen.CommentLike, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentLikeDo) FindInBatches(result *[]*orm_gen.CommentLike, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentLikeDo) Attrs(attrs ...field.AssignExpr) ICommentLikeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentLikeDo) Assign(attrs ...field.AssignExpr) ICommentLikeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentLikeDo) Joins(fields ...field.RelationField) ICommentLikeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentLikeDo) Preload(fields ...field.RelationField) ICommentLikeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentLikeDo) FirstOrInit() (*orm_gen.CommentLike, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentLike), nil
	}
}

func (c commentLikeDo) FirstOrCreate() (*orm_gen.CommentLike, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentLike), nil
	}
}

func (c commentLikeDo) FindByPage(offset int, limit int) (result []*orm_gen.CommentLike, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentLikeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentLikeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentLikeDo) Delete(models ...*orm_gen.CommentLike) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentLikeDo) withDO(do gen.Dao) *commentLikeDo {
	c.DO = *do.(*gen.DO)
	return c
}