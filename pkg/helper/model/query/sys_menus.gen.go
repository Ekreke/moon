// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/aide-family/moon/pkg/helper/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.ID = field.NewUint32(tableName, "id")
	_sysMenu.CreatedAt = field.NewField(tableName, "created_at")
	_sysMenu.UpdatedAt = field.NewField(tableName, "updated_at")
	_sysMenu.DeletedAt = field.NewUint(tableName, "deleted_at")
	_sysMenu.Name = field.NewString(tableName, "name")
	_sysMenu.EnName = field.NewString(tableName, "en_name")
	_sysMenu.Path = field.NewString(tableName, "path")
	_sysMenu.Status = field.NewInt(tableName, "status")
	_sysMenu.Icon = field.NewString(tableName, "icon")
	_sysMenu.ParentID = field.NewUint32(tableName, "parent_id")
	_sysMenu.Level = field.NewInt32(tableName, "level")
	_sysMenu.Parent = sysMenuBelongsToParent{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Parent", "model.SysMenu"),
		Parent: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Parent.Parent", "model.SysMenu"),
		},
	}

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo

	ALL       field.Asterisk
	ID        field.Uint32
	CreatedAt field.Field
	UpdatedAt field.Field
	DeletedAt field.Uint
	Name      field.String
	EnName    field.String
	Path      field.String
	Status    field.Int
	Icon      field.String
	ParentID  field.Uint32
	Level     field.Int32
	Parent    sysMenuBelongsToParent

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "id")
	s.CreatedAt = field.NewField(table, "created_at")
	s.UpdatedAt = field.NewField(table, "updated_at")
	s.DeletedAt = field.NewUint(table, "deleted_at")
	s.Name = field.NewString(table, "name")
	s.EnName = field.NewString(table, "en_name")
	s.Path = field.NewString(table, "path")
	s.Status = field.NewInt(table, "status")
	s.Icon = field.NewString(table, "icon")
	s.ParentID = field.NewUint32(table, "parent_id")
	s.Level = field.NewInt32(table, "level")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["en_name"] = s.EnName
	s.fieldMap["path"] = s.Path
	s.fieldMap["status"] = s.Status
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["level"] = s.Level

}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuBelongsToParent struct {
	db *gorm.DB

	field.RelationField

	Parent struct {
		field.RelationField
	}
}

func (a sysMenuBelongsToParent) Where(conds ...field.Expr) *sysMenuBelongsToParent {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a sysMenuBelongsToParent) WithContext(ctx context.Context) *sysMenuBelongsToParent {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sysMenuBelongsToParent) Session(session *gorm.Session) *sysMenuBelongsToParent {
	a.db = a.db.Session(session)
	return &a
}

func (a sysMenuBelongsToParent) Model(m *model.SysMenu) *sysMenuBelongsToParentTx {
	return &sysMenuBelongsToParentTx{a.db.Model(m).Association(a.Name())}
}

type sysMenuBelongsToParentTx struct{ tx *gorm.Association }

func (a sysMenuBelongsToParentTx) Find() (result *model.SysMenu, err error) {
	return result, a.tx.Find(&result)
}

func (a sysMenuBelongsToParentTx) Append(values ...*model.SysMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sysMenuBelongsToParentTx) Replace(values ...*model.SysMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sysMenuBelongsToParentTx) Delete(values ...*model.SysMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sysMenuBelongsToParentTx) Clear() error {
	return a.tx.Clear()
}

func (a sysMenuBelongsToParentTx) Count() int64 {
	return a.tx.Count()
}

type sysMenuDo struct{ gen.DO }

type ISysMenuDo interface {
	gen.SubQuery
	Debug() ISysMenuDo
	WithContext(ctx context.Context) ISysMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysMenuDo
	WriteDB() ISysMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysMenuDo
	Not(conds ...gen.Condition) ISysMenuDo
	Or(conds ...gen.Condition) ISysMenuDo
	Select(conds ...field.Expr) ISysMenuDo
	Where(conds ...gen.Condition) ISysMenuDo
	Order(conds ...field.Expr) ISysMenuDo
	Distinct(cols ...field.Expr) ISysMenuDo
	Omit(cols ...field.Expr) ISysMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	Group(cols ...field.Expr) ISysMenuDo
	Having(conds ...gen.Condition) ISysMenuDo
	Limit(limit int) ISysMenuDo
	Offset(offset int) ISysMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo
	Unscoped() ISysMenuDo
	Create(values ...*model.SysMenu) error
	CreateInBatches(values []*model.SysMenu, batchSize int) error
	Save(values ...*model.SysMenu) error
	First() (*model.SysMenu, error)
	Take() (*model.SysMenu, error)
	Last() (*model.SysMenu, error)
	Find() ([]*model.SysMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error)
	FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysMenuDo
	Assign(attrs ...field.AssignExpr) ISysMenuDo
	Joins(fields ...field.RelationField) ISysMenuDo
	Preload(fields ...field.RelationField) ISysMenuDo
	FirstOrInit() (*model.SysMenu, error)
	FirstOrCreate() (*model.SysMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysMenuDo) Debug() ISysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) ISysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() ISysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() ISysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) ISysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) ISysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) ISysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Order(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) ISysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) ISysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() ISysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*model.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
