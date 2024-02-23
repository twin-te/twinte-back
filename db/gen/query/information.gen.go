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

func newInformation(db *gorm.DB, opts ...gen.DOOption) information {
	_information := information{}

	_information.informationDo.UseDB(db, opts...)
	_information.informationDo.UseModel(&model.Information{})

	tableName := _information.informationDo.TableName()
	_information.ALL = field.NewAsterisk(tableName)
	_information.ID = field.NewString(tableName, "id")
	_information.Title = field.NewString(tableName, "title")
	_information.Content = field.NewString(tableName, "content")
	_information.PublishedAt = field.NewTime(tableName, "published_at")

	_information.fillFieldMap()

	return _information
}

type information struct {
	informationDo informationDo

	ALL         field.Asterisk
	ID          field.String
	Title       field.String
	Content     field.String
	PublishedAt field.Time

	fieldMap map[string]field.Expr
}

func (i information) Table(newTableName string) *information {
	i.informationDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i information) As(alias string) *information {
	i.informationDo.DO = *(i.informationDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *information) updateTableName(table string) *information {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewString(table, "id")
	i.Title = field.NewString(table, "title")
	i.Content = field.NewString(table, "content")
	i.PublishedAt = field.NewTime(table, "published_at")

	i.fillFieldMap()

	return i
}

func (i *information) WithContext(ctx context.Context) *informationDo {
	return i.informationDo.WithContext(ctx)
}

func (i information) TableName() string { return i.informationDo.TableName() }

func (i information) Alias() string { return i.informationDo.Alias() }

func (i information) Columns(cols ...field.Expr) gen.Columns { return i.informationDo.Columns(cols...) }

func (i *information) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *information) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 4)
	i.fieldMap["id"] = i.ID
	i.fieldMap["title"] = i.Title
	i.fieldMap["content"] = i.Content
	i.fieldMap["published_at"] = i.PublishedAt
}

func (i information) clone(db *gorm.DB) information {
	i.informationDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i information) replaceDB(db *gorm.DB) information {
	i.informationDo.ReplaceDB(db)
	return i
}

type informationDo struct{ gen.DO }

func (i informationDo) Debug() *informationDo {
	return i.withDO(i.DO.Debug())
}

func (i informationDo) WithContext(ctx context.Context) *informationDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i informationDo) ReadDB() *informationDo {
	return i.Clauses(dbresolver.Read)
}

func (i informationDo) WriteDB() *informationDo {
	return i.Clauses(dbresolver.Write)
}

func (i informationDo) Session(config *gorm.Session) *informationDo {
	return i.withDO(i.DO.Session(config))
}

func (i informationDo) Clauses(conds ...clause.Expression) *informationDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i informationDo) Returning(value interface{}, columns ...string) *informationDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i informationDo) Not(conds ...gen.Condition) *informationDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i informationDo) Or(conds ...gen.Condition) *informationDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i informationDo) Select(conds ...field.Expr) *informationDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i informationDo) Where(conds ...gen.Condition) *informationDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i informationDo) Order(conds ...field.Expr) *informationDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i informationDo) Distinct(cols ...field.Expr) *informationDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i informationDo) Omit(cols ...field.Expr) *informationDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i informationDo) Join(table schema.Tabler, on ...field.Expr) *informationDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i informationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *informationDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i informationDo) RightJoin(table schema.Tabler, on ...field.Expr) *informationDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i informationDo) Group(cols ...field.Expr) *informationDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i informationDo) Having(conds ...gen.Condition) *informationDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i informationDo) Limit(limit int) *informationDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i informationDo) Offset(offset int) *informationDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i informationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *informationDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i informationDo) Unscoped() *informationDo {
	return i.withDO(i.DO.Unscoped())
}

func (i informationDo) Create(values ...*model.Information) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i informationDo) CreateInBatches(values []*model.Information, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i informationDo) Save(values ...*model.Information) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i informationDo) First() (*model.Information, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Information), nil
	}
}

func (i informationDo) Take() (*model.Information, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Information), nil
	}
}

func (i informationDo) Last() (*model.Information, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Information), nil
	}
}

func (i informationDo) Find() ([]*model.Information, error) {
	result, err := i.DO.Find()
	return result.([]*model.Information), err
}

func (i informationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Information, err error) {
	buf := make([]*model.Information, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i informationDo) FindInBatches(result *[]*model.Information, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i informationDo) Attrs(attrs ...field.AssignExpr) *informationDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i informationDo) Assign(attrs ...field.AssignExpr) *informationDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i informationDo) Joins(fields ...field.RelationField) *informationDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i informationDo) Preload(fields ...field.RelationField) *informationDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i informationDo) FirstOrInit() (*model.Information, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Information), nil
	}
}

func (i informationDo) FirstOrCreate() (*model.Information, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Information), nil
	}
}

func (i informationDo) FindByPage(offset int, limit int) (result []*model.Information, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i informationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i informationDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i informationDo) Delete(models ...*model.Information) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *informationDo) withDO(do gen.Dao) *informationDo {
	i.DO = *do.(*gen.DO)
	return i
}