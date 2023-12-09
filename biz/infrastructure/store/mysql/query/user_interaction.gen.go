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

func newUserInteraction(db *gorm.DB, opts ...gen.DOOption) userInteraction {
	_userInteraction := userInteraction{}

	_userInteraction.userInteractionDo.UseDB(db, opts...)
	_userInteraction.userInteractionDo.UseModel(&orm_gen.UserInteraction{})

	tableName := _userInteraction.userInteractionDo.TableName()
	_userInteraction.ALL = field.NewAsterisk(tableName)
	_userInteraction.ID = field.NewField(tableName, "id")
	_userInteraction.CreatedAt = field.NewTime(tableName, "created_at")
	_userInteraction.DeletedAt = field.NewField(tableName, "deleted_at")
	_userInteraction.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userInteraction.UserID = field.NewString(tableName, "user_id")
	_userInteraction.FollowedID = field.NewString(tableName, "followed_id")

	_userInteraction.fillFieldMap()

	return _userInteraction
}

// userInteraction 社交关系表
type userInteraction struct {
	userInteractionDo

	ALL        field.Asterisk
	ID         field.Field  // 自然主键
	CreatedAt  field.Time   // 创建时间
	DeletedAt  field.Field  // 删除时间（软删除）
	UpdatedAt  field.Time   // 更新时间
	UserID     field.String // subject
	FollowedID field.String // object

	fieldMap map[string]field.Expr
}

func (u userInteraction) Table(newTableName string) *userInteraction {
	u.userInteractionDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userInteraction) As(alias string) *userInteraction {
	u.userInteractionDo.DO = *(u.userInteractionDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userInteraction) updateTableName(table string) *userInteraction {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewField(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.UserID = field.NewString(table, "user_id")
	u.FollowedID = field.NewString(table, "followed_id")

	u.fillFieldMap()

	return u
}

func (u *userInteraction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userInteraction) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["followed_id"] = u.FollowedID
}

func (u userInteraction) clone(db *gorm.DB) userInteraction {
	u.userInteractionDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userInteraction) replaceDB(db *gorm.DB) userInteraction {
	u.userInteractionDo.ReplaceDB(db)
	return u
}

type userInteractionDo struct{ gen.DO }

type IUserInteractionDo interface {
	gen.SubQuery
	Debug() IUserInteractionDo
	WithContext(ctx context.Context) IUserInteractionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserInteractionDo
	WriteDB() IUserInteractionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserInteractionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserInteractionDo
	Not(conds ...gen.Condition) IUserInteractionDo
	Or(conds ...gen.Condition) IUserInteractionDo
	Select(conds ...field.Expr) IUserInteractionDo
	Where(conds ...gen.Condition) IUserInteractionDo
	Order(conds ...field.Expr) IUserInteractionDo
	Distinct(cols ...field.Expr) IUserInteractionDo
	Omit(cols ...field.Expr) IUserInteractionDo
	Join(table schema.Tabler, on ...field.Expr) IUserInteractionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserInteractionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserInteractionDo
	Group(cols ...field.Expr) IUserInteractionDo
	Having(conds ...gen.Condition) IUserInteractionDo
	Limit(limit int) IUserInteractionDo
	Offset(offset int) IUserInteractionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInteractionDo
	Unscoped() IUserInteractionDo
	Create(values ...*orm_gen.UserInteraction) error
	CreateInBatches(values []*orm_gen.UserInteraction, batchSize int) error
	Save(values ...*orm_gen.UserInteraction) error
	First() (*orm_gen.UserInteraction, error)
	Take() (*orm_gen.UserInteraction, error)
	Last() (*orm_gen.UserInteraction, error)
	Find() ([]*orm_gen.UserInteraction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.UserInteraction, err error)
	FindInBatches(result *[]*orm_gen.UserInteraction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.UserInteraction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserInteractionDo
	Assign(attrs ...field.AssignExpr) IUserInteractionDo
	Joins(fields ...field.RelationField) IUserInteractionDo
	Preload(fields ...field.RelationField) IUserInteractionDo
	FirstOrInit() (*orm_gen.UserInteraction, error)
	FirstOrCreate() (*orm_gen.UserInteraction, error)
	FindByPage(offset int, limit int) (result []*orm_gen.UserInteraction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserInteractionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userInteractionDo) Debug() IUserInteractionDo {
	return u.withDO(u.DO.Debug())
}

func (u userInteractionDo) WithContext(ctx context.Context) IUserInteractionDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userInteractionDo) ReadDB() IUserInteractionDo {
	return u.Clauses(dbresolver.Read)
}

func (u userInteractionDo) WriteDB() IUserInteractionDo {
	return u.Clauses(dbresolver.Write)
}

func (u userInteractionDo) Session(config *gorm.Session) IUserInteractionDo {
	return u.withDO(u.DO.Session(config))
}

func (u userInteractionDo) Clauses(conds ...clause.Expression) IUserInteractionDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userInteractionDo) Returning(value interface{}, columns ...string) IUserInteractionDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userInteractionDo) Not(conds ...gen.Condition) IUserInteractionDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userInteractionDo) Or(conds ...gen.Condition) IUserInteractionDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userInteractionDo) Select(conds ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userInteractionDo) Where(conds ...gen.Condition) IUserInteractionDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userInteractionDo) Order(conds ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userInteractionDo) Distinct(cols ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userInteractionDo) Omit(cols ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userInteractionDo) Join(table schema.Tabler, on ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userInteractionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userInteractionDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userInteractionDo) Group(cols ...field.Expr) IUserInteractionDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userInteractionDo) Having(conds ...gen.Condition) IUserInteractionDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userInteractionDo) Limit(limit int) IUserInteractionDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userInteractionDo) Offset(offset int) IUserInteractionDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userInteractionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInteractionDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userInteractionDo) Unscoped() IUserInteractionDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userInteractionDo) Create(values ...*orm_gen.UserInteraction) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userInteractionDo) CreateInBatches(values []*orm_gen.UserInteraction, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userInteractionDo) Save(values ...*orm_gen.UserInteraction) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userInteractionDo) First() (*orm_gen.UserInteraction, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.UserInteraction), nil
	}
}

func (u userInteractionDo) Take() (*orm_gen.UserInteraction, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.UserInteraction), nil
	}
}

func (u userInteractionDo) Last() (*orm_gen.UserInteraction, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.UserInteraction), nil
	}
}

func (u userInteractionDo) Find() ([]*orm_gen.UserInteraction, error) {
	result, err := u.DO.Find()
	return result.([]*orm_gen.UserInteraction), err
}

func (u userInteractionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.UserInteraction, err error) {
	buf := make([]*orm_gen.UserInteraction, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userInteractionDo) FindInBatches(result *[]*orm_gen.UserInteraction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userInteractionDo) Attrs(attrs ...field.AssignExpr) IUserInteractionDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userInteractionDo) Assign(attrs ...field.AssignExpr) IUserInteractionDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userInteractionDo) Joins(fields ...field.RelationField) IUserInteractionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userInteractionDo) Preload(fields ...field.RelationField) IUserInteractionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userInteractionDo) FirstOrInit() (*orm_gen.UserInteraction, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.UserInteraction), nil
	}
}

func (u userInteractionDo) FirstOrCreate() (*orm_gen.UserInteraction, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.UserInteraction), nil
	}
}

func (u userInteractionDo) FindByPage(offset int, limit int) (result []*orm_gen.UserInteraction, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userInteractionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userInteractionDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userInteractionDo) Delete(models ...*orm_gen.UserInteraction) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userInteractionDo) withDO(do gen.Dao) *userInteractionDo {
	u.DO = *do.(*gen.DO)
	return u
}
