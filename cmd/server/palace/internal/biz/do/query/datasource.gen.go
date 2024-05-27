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

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
)

func newDatasource(db *gorm.DB, opts ...gen.DOOption) datasource {
	_datasource := datasource{}

	_datasource.datasourceDo.UseDB(db, opts...)
	_datasource.datasourceDo.UseModel(&model.Datasource{})

	tableName := _datasource.datasourceDo.TableName()
	_datasource.ALL = field.NewAsterisk(tableName)
	_datasource.ID = field.NewUint32(tableName, "id")
	_datasource.Name = field.NewString(tableName, "name")
	_datasource.Category = field.NewInt32(tableName, "category")
	_datasource.AuthExt = field.NewString(tableName, "auth_ext")
	_datasource.TeamID = field.NewUint32(tableName, "team_id")
	_datasource.Endpoint = field.NewString(tableName, "endpoint")
	_datasource.Status = field.NewInt(tableName, "status")
	_datasource.CreatedAt = field.NewTime(tableName, "created_at")
	_datasource.UpdatedAt = field.NewTime(tableName, "updated_at")
	_datasource.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_datasource.fillFieldMap()

	return _datasource
}

type datasource struct {
	datasourceDo

	ALL       field.Asterisk
	ID        field.Uint32
	Name      field.String // 数据源名称
	Category  field.Int32  // 数据源类型
	AuthExt   field.String // 认证扩展参数
	TeamID    field.Uint32 // 所属团队，为0时为全局可用
	Endpoint  field.String // 数据源地址
	Status    field.Int    // 数据源状态
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Int64  // 删除时间

	fieldMap map[string]field.Expr
}

func (d datasource) Table(newTableName string) *datasource {
	d.datasourceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d datasource) As(alias string) *datasource {
	d.datasourceDo.DO = *(d.datasourceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *datasource) updateTableName(table string) *datasource {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint32(table, "id")
	d.Name = field.NewString(table, "name")
	d.Category = field.NewInt32(table, "category")
	d.AuthExt = field.NewString(table, "auth_ext")
	d.TeamID = field.NewUint32(table, "team_id")
	d.Endpoint = field.NewString(table, "endpoint")
	d.Status = field.NewInt(table, "status")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewInt64(table, "deleted_at")

	d.fillFieldMap()

	return d
}

func (d *datasource) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *datasource) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["category"] = d.Category
	d.fieldMap["auth_ext"] = d.AuthExt
	d.fieldMap["team_id"] = d.TeamID
	d.fieldMap["endpoint"] = d.Endpoint
	d.fieldMap["status"] = d.Status
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
}

func (d datasource) clone(db *gorm.DB) datasource {
	d.datasourceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d datasource) replaceDB(db *gorm.DB) datasource {
	d.datasourceDo.ReplaceDB(db)
	return d
}

type datasourceDo struct{ gen.DO }

type IDatasourceDo interface {
	gen.SubQuery
	Debug() IDatasourceDo
	WithContext(ctx context.Context) IDatasourceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDatasourceDo
	WriteDB() IDatasourceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDatasourceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDatasourceDo
	Not(conds ...gen.Condition) IDatasourceDo
	Or(conds ...gen.Condition) IDatasourceDo
	Select(conds ...field.Expr) IDatasourceDo
	Where(conds ...gen.Condition) IDatasourceDo
	Order(conds ...field.Expr) IDatasourceDo
	Distinct(cols ...field.Expr) IDatasourceDo
	Omit(cols ...field.Expr) IDatasourceDo
	Join(table schema.Tabler, on ...field.Expr) IDatasourceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDatasourceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDatasourceDo
	Group(cols ...field.Expr) IDatasourceDo
	Having(conds ...gen.Condition) IDatasourceDo
	Limit(limit int) IDatasourceDo
	Offset(offset int) IDatasourceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDatasourceDo
	Unscoped() IDatasourceDo
	Create(values ...*model.Datasource) error
	CreateInBatches(values []*model.Datasource, batchSize int) error
	Save(values ...*model.Datasource) error
	First() (*model.Datasource, error)
	Take() (*model.Datasource, error)
	Last() (*model.Datasource, error)
	Find() ([]*model.Datasource, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Datasource, err error)
	FindInBatches(result *[]*model.Datasource, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Datasource) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDatasourceDo
	Assign(attrs ...field.AssignExpr) IDatasourceDo
	Joins(fields ...field.RelationField) IDatasourceDo
	Preload(fields ...field.RelationField) IDatasourceDo
	FirstOrInit() (*model.Datasource, error)
	FirstOrCreate() (*model.Datasource, error)
	FindByPage(offset int, limit int) (result []*model.Datasource, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDatasourceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d datasourceDo) Debug() IDatasourceDo {
	return d.withDO(d.DO.Debug())
}

func (d datasourceDo) WithContext(ctx context.Context) IDatasourceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d datasourceDo) ReadDB() IDatasourceDo {
	return d.Clauses(dbresolver.Read)
}

func (d datasourceDo) WriteDB() IDatasourceDo {
	return d.Clauses(dbresolver.Write)
}

func (d datasourceDo) Session(config *gorm.Session) IDatasourceDo {
	return d.withDO(d.DO.Session(config))
}

func (d datasourceDo) Clauses(conds ...clause.Expression) IDatasourceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d datasourceDo) Returning(value interface{}, columns ...string) IDatasourceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d datasourceDo) Not(conds ...gen.Condition) IDatasourceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d datasourceDo) Or(conds ...gen.Condition) IDatasourceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d datasourceDo) Select(conds ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d datasourceDo) Where(conds ...gen.Condition) IDatasourceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d datasourceDo) Order(conds ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d datasourceDo) Distinct(cols ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d datasourceDo) Omit(cols ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d datasourceDo) Join(table schema.Tabler, on ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d datasourceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d datasourceDo) RightJoin(table schema.Tabler, on ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d datasourceDo) Group(cols ...field.Expr) IDatasourceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d datasourceDo) Having(conds ...gen.Condition) IDatasourceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d datasourceDo) Limit(limit int) IDatasourceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d datasourceDo) Offset(offset int) IDatasourceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d datasourceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDatasourceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d datasourceDo) Unscoped() IDatasourceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d datasourceDo) Create(values ...*model.Datasource) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d datasourceDo) CreateInBatches(values []*model.Datasource, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d datasourceDo) Save(values ...*model.Datasource) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d datasourceDo) First() (*model.Datasource, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Datasource), nil
	}
}

func (d datasourceDo) Take() (*model.Datasource, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Datasource), nil
	}
}

func (d datasourceDo) Last() (*model.Datasource, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Datasource), nil
	}
}

func (d datasourceDo) Find() ([]*model.Datasource, error) {
	result, err := d.DO.Find()
	return result.([]*model.Datasource), err
}

func (d datasourceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Datasource, err error) {
	buf := make([]*model.Datasource, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d datasourceDo) FindInBatches(result *[]*model.Datasource, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d datasourceDo) Attrs(attrs ...field.AssignExpr) IDatasourceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d datasourceDo) Assign(attrs ...field.AssignExpr) IDatasourceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d datasourceDo) Joins(fields ...field.RelationField) IDatasourceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d datasourceDo) Preload(fields ...field.RelationField) IDatasourceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d datasourceDo) FirstOrInit() (*model.Datasource, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Datasource), nil
	}
}

func (d datasourceDo) FirstOrCreate() (*model.Datasource, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Datasource), nil
	}
}

func (d datasourceDo) FindByPage(offset int, limit int) (result []*model.Datasource, count int64, err error) {
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

func (d datasourceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d datasourceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d datasourceDo) Delete(models ...*model.Datasource) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *datasourceDo) withDO(do gen.Dao) *datasourceDo {
	d.DO = *do.(*gen.DO)
	return d
}