// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cold-runner/skylark/internal/model"
)

func newTopicLike(db *gorm.DB, opts ...gen.DOOption) topicLike {
	_topicLike := topicLike{}

	_topicLike.topicLikeDo.UseDB(db, opts...)
	_topicLike.topicLikeDo.UseModel(&model.TopicLike{})

	tableName := _topicLike.topicLikeDo.TableName()
	_topicLike.ALL = field.NewAsterisk(tableName)
	_topicLike.ID = field.NewInt64(tableName, "id")
	_topicLike.CreatedAt = field.NewTime(tableName, "created_at")
	_topicLike.DeletedAt = field.NewField(tableName, "deleted_at")
	_topicLike.UpdatedAt = field.NewTime(tableName, "updated_at")
	_topicLike.TopicID = field.NewInt64(tableName, "topic_id")
	_topicLike.UserID = field.NewInt64(tableName, "user_id")

	_topicLike.fillFieldMap()

	return _topicLike
}

type topicLike struct {
	topicLikeDo

	ALL       field.Asterisk
	ID        field.Int64 // 自然主键
	CreatedAt field.Time  // 创建时间
	DeletedAt field.Field // 删除时间（软删除）
	UpdatedAt field.Time  // 更新时间
	TopicID   field.Int64 // 话题id（考虑性能，不加外键）
	UserID    field.Int64 // 点赞人（考虑性能，不加外键）

	fieldMap map[string]field.Expr
}

func (t topicLike) Table(newTableName string) *topicLike {
	t.topicLikeDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t topicLike) As(alias string) *topicLike {
	t.topicLikeDo.DO = *(t.topicLikeDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *topicLike) updateTableName(table string) *topicLike {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.TopicID = field.NewInt64(table, "topic_id")
	t.UserID = field.NewInt64(table, "user_id")

	t.fillFieldMap()

	return t
}

func (t *topicLike) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *topicLike) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["topic_id"] = t.TopicID
	t.fieldMap["user_id"] = t.UserID
}

func (t topicLike) clone(db *gorm.DB) topicLike {
	t.topicLikeDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t topicLike) replaceDB(db *gorm.DB) topicLike {
	t.topicLikeDo.ReplaceDB(db)
	return t
}

type topicLikeDo struct{ gen.DO }

type ITopicLikeDo interface {
	gen.SubQuery
	Debug() ITopicLikeDo
	WithContext(ctx context.Context) ITopicLikeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITopicLikeDo
	WriteDB() ITopicLikeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITopicLikeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITopicLikeDo
	Not(conds ...gen.Condition) ITopicLikeDo
	Or(conds ...gen.Condition) ITopicLikeDo
	Select(conds ...field.Expr) ITopicLikeDo
	Where(conds ...gen.Condition) ITopicLikeDo
	Order(conds ...field.Expr) ITopicLikeDo
	Distinct(cols ...field.Expr) ITopicLikeDo
	Omit(cols ...field.Expr) ITopicLikeDo
	Join(table schema.Tabler, on ...field.Expr) ITopicLikeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITopicLikeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITopicLikeDo
	Group(cols ...field.Expr) ITopicLikeDo
	Having(conds ...gen.Condition) ITopicLikeDo
	Limit(limit int) ITopicLikeDo
	Offset(offset int) ITopicLikeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITopicLikeDo
	Unscoped() ITopicLikeDo
	Create(values ...*model.TopicLike) error
	CreateInBatches(values []*model.TopicLike, batchSize int) error
	Save(values ...*model.TopicLike) error
	First() (*model.TopicLike, error)
	Take() (*model.TopicLike, error)
	Last() (*model.TopicLike, error)
	Find() ([]*model.TopicLike, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TopicLike, err error)
	FindInBatches(result *[]*model.TopicLike, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TopicLike) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITopicLikeDo
	Assign(attrs ...field.AssignExpr) ITopicLikeDo
	Joins(fields ...field.RelationField) ITopicLikeDo
	Preload(fields ...field.RelationField) ITopicLikeDo
	FirstOrInit() (*model.TopicLike, error)
	FirstOrCreate() (*model.TopicLike, error)
	FindByPage(offset int, limit int) (result []*model.TopicLike, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITopicLikeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t topicLikeDo) Debug() ITopicLikeDo {
	return t.withDO(t.DO.Debug())
}

func (t topicLikeDo) WithContext(ctx context.Context) ITopicLikeDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t topicLikeDo) ReadDB() ITopicLikeDo {
	return t.Clauses(dbresolver.Read)
}

func (t topicLikeDo) WriteDB() ITopicLikeDo {
	return t.Clauses(dbresolver.Write)
}

func (t topicLikeDo) Session(config *gorm.Session) ITopicLikeDo {
	return t.withDO(t.DO.Session(config))
}

func (t topicLikeDo) Clauses(conds ...clause.Expression) ITopicLikeDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t topicLikeDo) Returning(value interface{}, columns ...string) ITopicLikeDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t topicLikeDo) Not(conds ...gen.Condition) ITopicLikeDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t topicLikeDo) Or(conds ...gen.Condition) ITopicLikeDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t topicLikeDo) Select(conds ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t topicLikeDo) Where(conds ...gen.Condition) ITopicLikeDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t topicLikeDo) Order(conds ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t topicLikeDo) Distinct(cols ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t topicLikeDo) Omit(cols ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t topicLikeDo) Join(table schema.Tabler, on ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t topicLikeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t topicLikeDo) RightJoin(table schema.Tabler, on ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t topicLikeDo) Group(cols ...field.Expr) ITopicLikeDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t topicLikeDo) Having(conds ...gen.Condition) ITopicLikeDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t topicLikeDo) Limit(limit int) ITopicLikeDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t topicLikeDo) Offset(offset int) ITopicLikeDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t topicLikeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITopicLikeDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t topicLikeDo) Unscoped() ITopicLikeDo {
	return t.withDO(t.DO.Unscoped())
}

func (t topicLikeDo) Create(values ...*model.TopicLike) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t topicLikeDo) CreateInBatches(values []*model.TopicLike, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t topicLikeDo) Save(values ...*model.TopicLike) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t topicLikeDo) First() (*model.TopicLike, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopicLike), nil
	}
}

func (t topicLikeDo) Take() (*model.TopicLike, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopicLike), nil
	}
}

func (t topicLikeDo) Last() (*model.TopicLike, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopicLike), nil
	}
}

func (t topicLikeDo) Find() ([]*model.TopicLike, error) {
	result, err := t.DO.Find()
	return result.([]*model.TopicLike), err
}

func (t topicLikeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TopicLike, err error) {
	buf := make([]*model.TopicLike, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t topicLikeDo) FindInBatches(result *[]*model.TopicLike, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t topicLikeDo) Attrs(attrs ...field.AssignExpr) ITopicLikeDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t topicLikeDo) Assign(attrs ...field.AssignExpr) ITopicLikeDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t topicLikeDo) Joins(fields ...field.RelationField) ITopicLikeDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t topicLikeDo) Preload(fields ...field.RelationField) ITopicLikeDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t topicLikeDo) FirstOrInit() (*model.TopicLike, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopicLike), nil
	}
}

func (t topicLikeDo) FirstOrCreate() (*model.TopicLike, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopicLike), nil
	}
}

func (t topicLikeDo) FindByPage(offset int, limit int) (result []*model.TopicLike, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t topicLikeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t topicLikeDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t topicLikeDo) Delete(models ...*model.TopicLike) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *topicLikeDo) withDO(do gen.Dao) *topicLikeDo {
	t.DO = *do.(*gen.DO)
	return t
}
