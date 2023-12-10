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

func newCategorie(db *gorm.DB, opts ...gen.DOOption) categorie {
	_categorie := categorie{}

	_categorie.categorieDo.UseDB(db, opts...)
	_categorie.categorieDo.UseModel(&orm_gen.Categorie{})

	tableName := _categorie.categorieDo.TableName()
	_categorie.ALL = field.NewAsterisk(tableName)
	_categorie.ID = field.NewField(tableName, "id")
	_categorie.CreatedAt = field.NewTime(tableName, "created_at")
	_categorie.DeletedAt = field.NewField(tableName, "deleted_at")
	_categorie.UpdatedAt = field.NewTime(tableName, "updated_at")
	_categorie.Name = field.NewString(tableName, "name")
	_categorie.BackgroundURL = field.NewString(tableName, "background_url")
	_categorie.Rank = field.NewInt64(tableName, "rank")
	_categorie.PlateID = field.NewString(tableName, "plate_id")
	_categorie.URL = field.NewString(tableName, "url")
	_categorie.Icon = field.NewString(tableName, "icon")

	_categorie.fillFieldMap()

	return _categorie
}

// categorie 归档表（板块）
type categorie struct {
	categorieDo

	ALL           field.Asterisk
	ID            field.Field  // 自然主键
	CreatedAt     field.Time   // 创建时间
	DeletedAt     field.Field  // 删除时间（软删除）
	UpdatedAt     field.Time   // 更新时间
	Name          field.String // 归档名称
	BackgroundURL field.String // 背景图片url
	Rank          field.Int64  // 板块排序权重
	PlateID       field.String // 板块id
	URL           field.String // 跳转url地址
	Icon          field.String // icon图标

	fieldMap map[string]field.Expr
}

func (c categorie) Table(newTableName string) *categorie {
	c.categorieDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categorie) As(alias string) *categorie {
	c.categorieDo.DO = *(c.categorieDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categorie) updateTableName(table string) *categorie {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.Name = field.NewString(table, "name")
	c.BackgroundURL = field.NewString(table, "background_url")
	c.Rank = field.NewInt64(table, "rank")
	c.PlateID = field.NewString(table, "plate_id")
	c.URL = field.NewString(table, "url")
	c.Icon = field.NewString(table, "icon")

	c.fillFieldMap()

	return c
}

func (c *categorie) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categorie) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["background_url"] = c.BackgroundURL
	c.fieldMap["rank"] = c.Rank
	c.fieldMap["plate_id"] = c.PlateID
	c.fieldMap["url"] = c.URL
	c.fieldMap["icon"] = c.Icon
}

func (c categorie) clone(db *gorm.DB) categorie {
	c.categorieDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categorie) replaceDB(db *gorm.DB) categorie {
	c.categorieDo.ReplaceDB(db)
	return c
}

type categorieDo struct{ gen.DO }

type ICategorieDo interface {
	gen.SubQuery
	Debug() ICategorieDo
	WithContext(ctx context.Context) ICategorieDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICategorieDo
	WriteDB() ICategorieDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICategorieDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICategorieDo
	Not(conds ...gen.Condition) ICategorieDo
	Or(conds ...gen.Condition) ICategorieDo
	Select(conds ...field.Expr) ICategorieDo
	Where(conds ...gen.Condition) ICategorieDo
	Order(conds ...field.Expr) ICategorieDo
	Distinct(cols ...field.Expr) ICategorieDo
	Omit(cols ...field.Expr) ICategorieDo
	Join(table schema.Tabler, on ...field.Expr) ICategorieDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICategorieDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICategorieDo
	Group(cols ...field.Expr) ICategorieDo
	Having(conds ...gen.Condition) ICategorieDo
	Limit(limit int) ICategorieDo
	Offset(offset int) ICategorieDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICategorieDo
	Unscoped() ICategorieDo
	Create(values ...*orm_gen.Categorie) error
	CreateInBatches(values []*orm_gen.Categorie, batchSize int) error
	Save(values ...*orm_gen.Categorie) error
	First() (*orm_gen.Categorie, error)
	Take() (*orm_gen.Categorie, error)
	Last() (*orm_gen.Categorie, error)
	Find() ([]*orm_gen.Categorie, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Categorie, err error)
	FindInBatches(result *[]*orm_gen.Categorie, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.Categorie) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICategorieDo
	Assign(attrs ...field.AssignExpr) ICategorieDo
	Joins(fields ...field.RelationField) ICategorieDo
	Preload(fields ...field.RelationField) ICategorieDo
	FirstOrInit() (*orm_gen.Categorie, error)
	FirstOrCreate() (*orm_gen.Categorie, error)
	FindByPage(offset int, limit int) (result []*orm_gen.Categorie, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICategorieDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c categorieDo) Debug() ICategorieDo {
	return c.withDO(c.DO.Debug())
}

func (c categorieDo) WithContext(ctx context.Context) ICategorieDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categorieDo) ReadDB() ICategorieDo {
	return c.Clauses(dbresolver.Read)
}

func (c categorieDo) WriteDB() ICategorieDo {
	return c.Clauses(dbresolver.Write)
}

func (c categorieDo) Session(config *gorm.Session) ICategorieDo {
	return c.withDO(c.DO.Session(config))
}

func (c categorieDo) Clauses(conds ...clause.Expression) ICategorieDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categorieDo) Returning(value interface{}, columns ...string) ICategorieDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categorieDo) Not(conds ...gen.Condition) ICategorieDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categorieDo) Or(conds ...gen.Condition) ICategorieDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categorieDo) Select(conds ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categorieDo) Where(conds ...gen.Condition) ICategorieDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categorieDo) Order(conds ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categorieDo) Distinct(cols ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categorieDo) Omit(cols ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categorieDo) Join(table schema.Tabler, on ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categorieDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categorieDo) RightJoin(table schema.Tabler, on ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categorieDo) Group(cols ...field.Expr) ICategorieDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categorieDo) Having(conds ...gen.Condition) ICategorieDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categorieDo) Limit(limit int) ICategorieDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categorieDo) Offset(offset int) ICategorieDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categorieDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICategorieDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categorieDo) Unscoped() ICategorieDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categorieDo) Create(values ...*orm_gen.Categorie) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categorieDo) CreateInBatches(values []*orm_gen.Categorie, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categorieDo) Save(values ...*orm_gen.Categorie) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categorieDo) First() (*orm_gen.Categorie, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Categorie), nil
	}
}

func (c categorieDo) Take() (*orm_gen.Categorie, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Categorie), nil
	}
}

func (c categorieDo) Last() (*orm_gen.Categorie, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Categorie), nil
	}
}

func (c categorieDo) Find() ([]*orm_gen.Categorie, error) {
	result, err := c.DO.Find()
	return result.([]*orm_gen.Categorie), err
}

func (c categorieDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.Categorie, err error) {
	buf := make([]*orm_gen.Categorie, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categorieDo) FindInBatches(result *[]*orm_gen.Categorie, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categorieDo) Attrs(attrs ...field.AssignExpr) ICategorieDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categorieDo) Assign(attrs ...field.AssignExpr) ICategorieDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categorieDo) Joins(fields ...field.RelationField) ICategorieDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categorieDo) Preload(fields ...field.RelationField) ICategorieDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categorieDo) FirstOrInit() (*orm_gen.Categorie, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Categorie), nil
	}
}

func (c categorieDo) FirstOrCreate() (*orm_gen.Categorie, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.Categorie), nil
	}
}

func (c categorieDo) FindByPage(offset int, limit int) (result []*orm_gen.Categorie, count int64, err error) {
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

func (c categorieDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categorieDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categorieDo) Delete(models ...*orm_gen.Categorie) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categorieDo) withDO(do gen.Dao) *categorieDo {
	c.DO = *do.(*gen.DO)
	return c
}
