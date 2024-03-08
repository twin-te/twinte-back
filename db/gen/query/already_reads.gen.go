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

	"github.com/twin-te/twinte-back/db/gen/model"
)

func newAlreadyRead(db *gorm.DB, opts ...gen.DOOption) alreadyRead {
	_alreadyRead := alreadyRead{}

	_alreadyRead.alreadyReadDo.UseDB(db, opts...)
	_alreadyRead.alreadyReadDo.UseModel(&model.AlreadyRead{})

	tableName := _alreadyRead.alreadyReadDo.TableName()
	_alreadyRead.ALL = field.NewAsterisk(tableName)
	_alreadyRead.ID = field.NewString(tableName, "id")
	_alreadyRead.InformationID = field.NewString(tableName, "information_id")
	_alreadyRead.ReadUser = field.NewString(tableName, "read_user")
	_alreadyRead.ReadAt = field.NewTime(tableName, "read_at")

	_alreadyRead.fillFieldMap()

	return _alreadyRead
}

type alreadyRead struct {
	alreadyReadDo alreadyReadDo

	ALL           field.Asterisk
	ID            field.String
	InformationID field.String
	ReadUser      field.String
	ReadAt        field.Time

	fieldMap map[string]field.Expr
}

func (a alreadyRead) Table(newTableName string) *alreadyRead {
	a.alreadyReadDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alreadyRead) As(alias string) *alreadyRead {
	a.alreadyReadDo.DO = *(a.alreadyReadDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alreadyRead) updateTableName(table string) *alreadyRead {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.InformationID = field.NewString(table, "information_id")
	a.ReadUser = field.NewString(table, "read_user")
	a.ReadAt = field.NewTime(table, "read_at")

	a.fillFieldMap()

	return a
}

func (a *alreadyRead) WithContext(ctx context.Context) *alreadyReadDo {
	return a.alreadyReadDo.WithContext(ctx)
}

func (a alreadyRead) TableName() string { return a.alreadyReadDo.TableName() }

func (a alreadyRead) Alias() string { return a.alreadyReadDo.Alias() }

func (a alreadyRead) Columns(cols ...field.Expr) gen.Columns { return a.alreadyReadDo.Columns(cols...) }

func (a *alreadyRead) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alreadyRead) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["information_id"] = a.InformationID
	a.fieldMap["read_user"] = a.ReadUser
	a.fieldMap["read_at"] = a.ReadAt
}

func (a alreadyRead) clone(db *gorm.DB) alreadyRead {
	a.alreadyReadDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alreadyRead) replaceDB(db *gorm.DB) alreadyRead {
	a.alreadyReadDo.ReplaceDB(db)
	return a
}

type alreadyReadDo struct{ gen.DO }

func (a alreadyReadDo) Debug() *alreadyReadDo {
	return a.withDO(a.DO.Debug())
}

func (a alreadyReadDo) WithContext(ctx context.Context) *alreadyReadDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alreadyReadDo) ReadDB() *alreadyReadDo {
	return a.Clauses(dbresolver.Read)
}

func (a alreadyReadDo) WriteDB() *alreadyReadDo {
	return a.Clauses(dbresolver.Write)
}

func (a alreadyReadDo) Session(config *gorm.Session) *alreadyReadDo {
	return a.withDO(a.DO.Session(config))
}

func (a alreadyReadDo) Clauses(conds ...clause.Expression) *alreadyReadDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alreadyReadDo) Returning(value interface{}, columns ...string) *alreadyReadDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alreadyReadDo) Not(conds ...gen.Condition) *alreadyReadDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alreadyReadDo) Or(conds ...gen.Condition) *alreadyReadDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alreadyReadDo) Select(conds ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alreadyReadDo) Where(conds ...gen.Condition) *alreadyReadDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alreadyReadDo) Order(conds ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alreadyReadDo) Distinct(cols ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alreadyReadDo) Omit(cols ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alreadyReadDo) Join(table schema.Tabler, on ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alreadyReadDo) LeftJoin(table schema.Tabler, on ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alreadyReadDo) RightJoin(table schema.Tabler, on ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alreadyReadDo) Group(cols ...field.Expr) *alreadyReadDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alreadyReadDo) Having(conds ...gen.Condition) *alreadyReadDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alreadyReadDo) Limit(limit int) *alreadyReadDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alreadyReadDo) Offset(offset int) *alreadyReadDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alreadyReadDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *alreadyReadDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alreadyReadDo) Unscoped() *alreadyReadDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alreadyReadDo) Create(values ...*model.AlreadyRead) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alreadyReadDo) CreateInBatches(values []*model.AlreadyRead, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alreadyReadDo) Save(values ...*model.AlreadyRead) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alreadyReadDo) First() (*model.AlreadyRead, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlreadyRead), nil
	}
}

func (a alreadyReadDo) Take() (*model.AlreadyRead, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlreadyRead), nil
	}
}

func (a alreadyReadDo) Last() (*model.AlreadyRead, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlreadyRead), nil
	}
}

func (a alreadyReadDo) Find() ([]*model.AlreadyRead, error) {
	result, err := a.DO.Find()
	return result.([]*model.AlreadyRead), err
}

func (a alreadyReadDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlreadyRead, err error) {
	buf := make([]*model.AlreadyRead, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alreadyReadDo) FindInBatches(result *[]*model.AlreadyRead, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alreadyReadDo) Attrs(attrs ...field.AssignExpr) *alreadyReadDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alreadyReadDo) Assign(attrs ...field.AssignExpr) *alreadyReadDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alreadyReadDo) Joins(fields ...field.RelationField) *alreadyReadDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alreadyReadDo) Preload(fields ...field.RelationField) *alreadyReadDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alreadyReadDo) FirstOrInit() (*model.AlreadyRead, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlreadyRead), nil
	}
}

func (a alreadyReadDo) FirstOrCreate() (*model.AlreadyRead, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlreadyRead), nil
	}
}

func (a alreadyReadDo) FindByPage(offset int, limit int) (result []*model.AlreadyRead, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a alreadyReadDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alreadyReadDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alreadyReadDo) Delete(models ...*model.AlreadyRead) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alreadyReadDo) withDO(do gen.Dao) *alreadyReadDo {
	a.DO = *do.(*gen.DO)
	return a
}