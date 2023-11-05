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

	"github.com/cold-runner/skylark/internal/model"
)

func newDraft(db *gorm.DB, opts ...gen.DOOption) draft {
	_draft := draft{}

	_draft.draftDo.UseDB(db, opts...)
	_draft.draftDo.UseModel(&model.Draft{})

	tableName := _draft.draftDo.TableName()
	_draft.ALL = field.NewAsterisk(tableName)
	_draft.ID = field.NewInt64(tableName, "id")
	_draft.CreatedAt = field.NewTime(tableName, "created_at")
	_draft.UpdatedAt = field.NewTime(tableName, "updated_at")
	_draft.DeletedAt = field.NewField(tableName, "deleted_at")
	_draft.Title = field.NewString(tableName, "title")
	_draft.Content = field.NewString(tableName, "content")
	_draft.UserID = field.NewInt64(tableName, "user_id")

	_draft.fillFieldMap()

	return _draft
}

type draft struct {
	draftDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Title     field.String
	Content   field.String
	UserID    field.Int64

	fieldMap map[string]field.Expr
}

func (d draft) Table(newTableName string) *draft {
	d.draftDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d draft) As(alias string) *draft {
	d.draftDo.DO = *(d.draftDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *draft) updateTableName(table string) *draft {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Title = field.NewString(table, "title")
	d.Content = field.NewString(table, "content")
	d.UserID = field.NewInt64(table, "user_id")

	d.fillFieldMap()

	return d
}

func (d *draft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *draft) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["title"] = d.Title
	d.fieldMap["content"] = d.Content
	d.fieldMap["user_id"] = d.UserID
}

func (d draft) clone(db *gorm.DB) draft {
	d.draftDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d draft) replaceDB(db *gorm.DB) draft {
	d.draftDo.ReplaceDB(db)
	return d
}

type draftDo struct{ gen.DO }

type IDraftDo interface {
	gen.SubQuery
	Debug() IDraftDo
	WithContext(ctx context.Context) IDraftDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDraftDo
	WriteDB() IDraftDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDraftDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDraftDo
	Not(conds ...gen.Condition) IDraftDo
	Or(conds ...gen.Condition) IDraftDo
	Select(conds ...field.Expr) IDraftDo
	Where(conds ...gen.Condition) IDraftDo
	Order(conds ...field.Expr) IDraftDo
	Distinct(cols ...field.Expr) IDraftDo
	Omit(cols ...field.Expr) IDraftDo
	Join(table schema.Tabler, on ...field.Expr) IDraftDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDraftDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDraftDo
	Group(cols ...field.Expr) IDraftDo
	Having(conds ...gen.Condition) IDraftDo
	Limit(limit int) IDraftDo
	Offset(offset int) IDraftDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDraftDo
	Unscoped() IDraftDo
	Create(values ...*model.Draft) error
	CreateInBatches(values []*model.Draft, batchSize int) error
	Save(values ...*model.Draft) error
	First() (*model.Draft, error)
	Take() (*model.Draft, error)
	Last() (*model.Draft, error)
	Find() ([]*model.Draft, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Draft, err error)
	FindInBatches(result *[]*model.Draft, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Draft) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDraftDo
	Assign(attrs ...field.AssignExpr) IDraftDo
	Joins(fields ...field.RelationField) IDraftDo
	Preload(fields ...field.RelationField) IDraftDo
	FirstOrInit() (*model.Draft, error)
	FirstOrCreate() (*model.Draft, error)
	FindByPage(offset int, limit int) (result []*model.Draft, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDraftDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d draftDo) Debug() IDraftDo {
	return d.withDO(d.DO.Debug())
}

func (d draftDo) WithContext(ctx context.Context) IDraftDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d draftDo) ReadDB() IDraftDo {
	return d.Clauses(dbresolver.Read)
}

func (d draftDo) WriteDB() IDraftDo {
	return d.Clauses(dbresolver.Write)
}

func (d draftDo) Session(config *gorm.Session) IDraftDo {
	return d.withDO(d.DO.Session(config))
}

func (d draftDo) Clauses(conds ...clause.Expression) IDraftDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d draftDo) Returning(value interface{}, columns ...string) IDraftDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d draftDo) Not(conds ...gen.Condition) IDraftDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d draftDo) Or(conds ...gen.Condition) IDraftDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d draftDo) Select(conds ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d draftDo) Where(conds ...gen.Condition) IDraftDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d draftDo) Order(conds ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d draftDo) Distinct(cols ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d draftDo) Omit(cols ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d draftDo) Join(table schema.Tabler, on ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d draftDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDraftDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d draftDo) RightJoin(table schema.Tabler, on ...field.Expr) IDraftDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d draftDo) Group(cols ...field.Expr) IDraftDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d draftDo) Having(conds ...gen.Condition) IDraftDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d draftDo) Limit(limit int) IDraftDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d draftDo) Offset(offset int) IDraftDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d draftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDraftDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d draftDo) Unscoped() IDraftDo {
	return d.withDO(d.DO.Unscoped())
}

func (d draftDo) Create(values ...*model.Draft) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d draftDo) CreateInBatches(values []*model.Draft, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d draftDo) Save(values ...*model.Draft) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d draftDo) First() (*model.Draft, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Draft), nil
	}
}

func (d draftDo) Take() (*model.Draft, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Draft), nil
	}
}

func (d draftDo) Last() (*model.Draft, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Draft), nil
	}
}

func (d draftDo) Find() ([]*model.Draft, error) {
	result, err := d.DO.Find()
	return result.([]*model.Draft), err
}

func (d draftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Draft, err error) {
	buf := make([]*model.Draft, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d draftDo) FindInBatches(result *[]*model.Draft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d draftDo) Attrs(attrs ...field.AssignExpr) IDraftDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d draftDo) Assign(attrs ...field.AssignExpr) IDraftDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d draftDo) Joins(fields ...field.RelationField) IDraftDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d draftDo) Preload(fields ...field.RelationField) IDraftDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d draftDo) FirstOrInit() (*model.Draft, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Draft), nil
	}
}

func (d draftDo) FirstOrCreate() (*model.Draft, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Draft), nil
	}
}

func (d draftDo) FindByPage(offset int, limit int) (result []*model.Draft, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d draftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d draftDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d draftDo) Delete(models ...*model.Draft) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *draftDo) withDO(do gen.Dao) *draftDo {
	d.DO = *do.(*gen.DO)
	return d
}
