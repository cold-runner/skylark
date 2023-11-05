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

func newPlate(db *gorm.DB, opts ...gen.DOOption) plate {
	_plate := plate{}

	_plate.plateDo.UseDB(db, opts...)
	_plate.plateDo.UseModel(&model.Plate{})

	tableName := _plate.plateDo.TableName()
	_plate.ALL = field.NewAsterisk(tableName)
	_plate.ID = field.NewInt64(tableName, "id")
	_plate.CreatedAt = field.NewTime(tableName, "created_at")
	_plate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_plate.DeletedAt = field.NewField(tableName, "deleted_at")
	_plate.Name = field.NewString(tableName, "name")

	_plate.fillFieldMap()

	return _plate
}

type plate struct {
	plateDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String

	fieldMap map[string]field.Expr
}

func (p plate) Table(newTableName string) *plate {
	p.plateDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plate) As(alias string) *plate {
	p.plateDo.DO = *(p.plateDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plate) updateTableName(table string) *plate {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Name = field.NewString(table, "name")

	p.fillFieldMap()

	return p
}

func (p *plate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plate) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["name"] = p.Name
}

func (p plate) clone(db *gorm.DB) plate {
	p.plateDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plate) replaceDB(db *gorm.DB) plate {
	p.plateDo.ReplaceDB(db)
	return p
}

type plateDo struct{ gen.DO }

type IPlateDo interface {
	gen.SubQuery
	Debug() IPlateDo
	WithContext(ctx context.Context) IPlateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlateDo
	WriteDB() IPlateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlateDo
	Not(conds ...gen.Condition) IPlateDo
	Or(conds ...gen.Condition) IPlateDo
	Select(conds ...field.Expr) IPlateDo
	Where(conds ...gen.Condition) IPlateDo
	Order(conds ...field.Expr) IPlateDo
	Distinct(cols ...field.Expr) IPlateDo
	Omit(cols ...field.Expr) IPlateDo
	Join(table schema.Tabler, on ...field.Expr) IPlateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlateDo
	Group(cols ...field.Expr) IPlateDo
	Having(conds ...gen.Condition) IPlateDo
	Limit(limit int) IPlateDo
	Offset(offset int) IPlateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlateDo
	Unscoped() IPlateDo
	Create(values ...*model.Plate) error
	CreateInBatches(values []*model.Plate, batchSize int) error
	Save(values ...*model.Plate) error
	First() (*model.Plate, error)
	Take() (*model.Plate, error)
	Last() (*model.Plate, error)
	Find() ([]*model.Plate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Plate, err error)
	FindInBatches(result *[]*model.Plate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Plate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlateDo
	Assign(attrs ...field.AssignExpr) IPlateDo
	Joins(fields ...field.RelationField) IPlateDo
	Preload(fields ...field.RelationField) IPlateDo
	FirstOrInit() (*model.Plate, error)
	FirstOrCreate() (*model.Plate, error)
	FindByPage(offset int, limit int) (result []*model.Plate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p plateDo) Debug() IPlateDo {
	return p.withDO(p.DO.Debug())
}

func (p plateDo) WithContext(ctx context.Context) IPlateDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p plateDo) ReadDB() IPlateDo {
	return p.Clauses(dbresolver.Read)
}

func (p plateDo) WriteDB() IPlateDo {
	return p.Clauses(dbresolver.Write)
}

func (p plateDo) Session(config *gorm.Session) IPlateDo {
	return p.withDO(p.DO.Session(config))
}

func (p plateDo) Clauses(conds ...clause.Expression) IPlateDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p plateDo) Returning(value interface{}, columns ...string) IPlateDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p plateDo) Not(conds ...gen.Condition) IPlateDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p plateDo) Or(conds ...gen.Condition) IPlateDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p plateDo) Select(conds ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p plateDo) Where(conds ...gen.Condition) IPlateDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p plateDo) Order(conds ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p plateDo) Distinct(cols ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p plateDo) Omit(cols ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p plateDo) Join(table schema.Tabler, on ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p plateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlateDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p plateDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlateDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p plateDo) Group(cols ...field.Expr) IPlateDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p plateDo) Having(conds ...gen.Condition) IPlateDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p plateDo) Limit(limit int) IPlateDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p plateDo) Offset(offset int) IPlateDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p plateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlateDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p plateDo) Unscoped() IPlateDo {
	return p.withDO(p.DO.Unscoped())
}

func (p plateDo) Create(values ...*model.Plate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p plateDo) CreateInBatches(values []*model.Plate, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p plateDo) Save(values ...*model.Plate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p plateDo) First() (*model.Plate, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plate), nil
	}
}

func (p plateDo) Take() (*model.Plate, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plate), nil
	}
}

func (p plateDo) Last() (*model.Plate, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plate), nil
	}
}

func (p plateDo) Find() ([]*model.Plate, error) {
	result, err := p.DO.Find()
	return result.([]*model.Plate), err
}

func (p plateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Plate, err error) {
	buf := make([]*model.Plate, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p plateDo) FindInBatches(result *[]*model.Plate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p plateDo) Attrs(attrs ...field.AssignExpr) IPlateDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p plateDo) Assign(attrs ...field.AssignExpr) IPlateDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p plateDo) Joins(fields ...field.RelationField) IPlateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p plateDo) Preload(fields ...field.RelationField) IPlateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p plateDo) FirstOrInit() (*model.Plate, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plate), nil
	}
}

func (p plateDo) FirstOrCreate() (*model.Plate, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plate), nil
	}
}

func (p plateDo) FindByPage(offset int, limit int) (result []*model.Plate, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p plateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p plateDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p plateDo) Delete(models ...*model.Plate) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *plateDo) withDO(do gen.Dao) *plateDo {
	p.DO = *do.(*gen.DO)
	return p
}
