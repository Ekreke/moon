// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package bizquery

import (
	"context"

	"github.com/aide-family/moon/pkg/helper/model/bizmodel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newSysTeamMenu(db *gorm.DB, opts ...gen.DOOption) sysTeamMenu {
	_sysTeamMenu := sysTeamMenu{}

	_sysTeamMenu.sysTeamMenuDo.UseDB(db, opts...)
	_sysTeamMenu.sysTeamMenuDo.UseModel(&bizmodel.SysTeamMenu{})

	tableName := _sysTeamMenu.sysTeamMenuDo.TableName()
	_sysTeamMenu.ALL = field.NewAsterisk(tableName)
	_sysTeamMenu.ID = field.NewUint32(tableName, "id")
	_sysTeamMenu.CreatedAt = field.NewField(tableName, "created_at")
	_sysTeamMenu.UpdatedAt = field.NewField(tableName, "updated_at")
	_sysTeamMenu.DeletedAt = field.NewUint(tableName, "deleted_at")
	_sysTeamMenu.Name = field.NewString(tableName, "name")
	_sysTeamMenu.EnName = field.NewString(tableName, "en_name")
	_sysTeamMenu.Path = field.NewString(tableName, "path")
	_sysTeamMenu.Status = field.NewInt(tableName, "status")
	_sysTeamMenu.Icon = field.NewString(tableName, "icon")
	_sysTeamMenu.ParentID = field.NewUint32(tableName, "parent_id")
	_sysTeamMenu.Level = field.NewInt32(tableName, "level")
	_sysTeamMenu.Parent = sysTeamMenuBelongsToParent{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Parent", "bizmodel.SysTeamMenu"),
		Parent: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Parent.Parent", "bizmodel.SysTeamMenu"),
		},
	}

	_sysTeamMenu.fillFieldMap()

	return _sysTeamMenu
}

type sysTeamMenu struct {
	sysTeamMenuDo

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
	Parent    sysTeamMenuBelongsToParent

	fieldMap map[string]field.Expr
}

func (s sysTeamMenu) Table(newTableName string) *sysTeamMenu {
	s.sysTeamMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTeamMenu) As(alias string) *sysTeamMenu {
	s.sysTeamMenuDo.DO = *(s.sysTeamMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTeamMenu) updateTableName(table string) *sysTeamMenu {
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

func (s *sysTeamMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTeamMenu) fillFieldMap() {
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

func (s sysTeamMenu) clone(db *gorm.DB) sysTeamMenu {
	s.sysTeamMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTeamMenu) replaceDB(db *gorm.DB) sysTeamMenu {
	s.sysTeamMenuDo.ReplaceDB(db)
	return s
}

type sysTeamMenuBelongsToParent struct {
	db *gorm.DB

	field.RelationField

	Parent struct {
		field.RelationField
	}
}

func (a sysTeamMenuBelongsToParent) Where(conds ...field.Expr) *sysTeamMenuBelongsToParent {
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

func (a sysTeamMenuBelongsToParent) WithContext(ctx context.Context) *sysTeamMenuBelongsToParent {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sysTeamMenuBelongsToParent) Session(session *gorm.Session) *sysTeamMenuBelongsToParent {
	a.db = a.db.Session(session)
	return &a
}

func (a sysTeamMenuBelongsToParent) Model(m *bizmodel.SysTeamMenu) *sysTeamMenuBelongsToParentTx {
	return &sysTeamMenuBelongsToParentTx{a.db.Model(m).Association(a.Name())}
}

type sysTeamMenuBelongsToParentTx struct{ tx *gorm.Association }

func (a sysTeamMenuBelongsToParentTx) Find() (result *bizmodel.SysTeamMenu, err error) {
	return result, a.tx.Find(&result)
}

func (a sysTeamMenuBelongsToParentTx) Append(values ...*bizmodel.SysTeamMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sysTeamMenuBelongsToParentTx) Replace(values ...*bizmodel.SysTeamMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sysTeamMenuBelongsToParentTx) Delete(values ...*bizmodel.SysTeamMenu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sysTeamMenuBelongsToParentTx) Clear() error {
	return a.tx.Clear()
}

func (a sysTeamMenuBelongsToParentTx) Count() int64 {
	return a.tx.Count()
}

type sysTeamMenuDo struct{ gen.DO }

type ISysTeamMenuDo interface {
	gen.SubQuery
	Debug() ISysTeamMenuDo
	WithContext(ctx context.Context) ISysTeamMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTeamMenuDo
	WriteDB() ISysTeamMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTeamMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTeamMenuDo
	Not(conds ...gen.Condition) ISysTeamMenuDo
	Or(conds ...gen.Condition) ISysTeamMenuDo
	Select(conds ...field.Expr) ISysTeamMenuDo
	Where(conds ...gen.Condition) ISysTeamMenuDo
	Order(conds ...field.Expr) ISysTeamMenuDo
	Distinct(cols ...field.Expr) ISysTeamMenuDo
	Omit(cols ...field.Expr) ISysTeamMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo
	Group(cols ...field.Expr) ISysTeamMenuDo
	Having(conds ...gen.Condition) ISysTeamMenuDo
	Limit(limit int) ISysTeamMenuDo
	Offset(offset int) ISysTeamMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamMenuDo
	Unscoped() ISysTeamMenuDo
	Create(values ...*bizmodel.SysTeamMenu) error
	CreateInBatches(values []*bizmodel.SysTeamMenu, batchSize int) error
	Save(values ...*bizmodel.SysTeamMenu) error
	First() (*bizmodel.SysTeamMenu, error)
	Take() (*bizmodel.SysTeamMenu, error)
	Last() (*bizmodel.SysTeamMenu, error)
	Find() ([]*bizmodel.SysTeamMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.SysTeamMenu, err error)
	FindInBatches(result *[]*bizmodel.SysTeamMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*bizmodel.SysTeamMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTeamMenuDo
	Assign(attrs ...field.AssignExpr) ISysTeamMenuDo
	Joins(fields ...field.RelationField) ISysTeamMenuDo
	Preload(fields ...field.RelationField) ISysTeamMenuDo
	FirstOrInit() (*bizmodel.SysTeamMenu, error)
	FirstOrCreate() (*bizmodel.SysTeamMenu, error)
	FindByPage(offset int, limit int) (result []*bizmodel.SysTeamMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTeamMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTeamMenuDo) Debug() ISysTeamMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTeamMenuDo) WithContext(ctx context.Context) ISysTeamMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTeamMenuDo) ReadDB() ISysTeamMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTeamMenuDo) WriteDB() ISysTeamMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTeamMenuDo) Session(config *gorm.Session) ISysTeamMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTeamMenuDo) Clauses(conds ...clause.Expression) ISysTeamMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTeamMenuDo) Returning(value interface{}, columns ...string) ISysTeamMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTeamMenuDo) Not(conds ...gen.Condition) ISysTeamMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTeamMenuDo) Or(conds ...gen.Condition) ISysTeamMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTeamMenuDo) Select(conds ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTeamMenuDo) Where(conds ...gen.Condition) ISysTeamMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTeamMenuDo) Order(conds ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTeamMenuDo) Distinct(cols ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTeamMenuDo) Omit(cols ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTeamMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTeamMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTeamMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTeamMenuDo) Group(cols ...field.Expr) ISysTeamMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTeamMenuDo) Having(conds ...gen.Condition) ISysTeamMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTeamMenuDo) Limit(limit int) ISysTeamMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTeamMenuDo) Offset(offset int) ISysTeamMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTeamMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTeamMenuDo) Unscoped() ISysTeamMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTeamMenuDo) Create(values ...*bizmodel.SysTeamMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTeamMenuDo) CreateInBatches(values []*bizmodel.SysTeamMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTeamMenuDo) Save(values ...*bizmodel.SysTeamMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTeamMenuDo) First() (*bizmodel.SysTeamMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamMenu), nil
	}
}

func (s sysTeamMenuDo) Take() (*bizmodel.SysTeamMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamMenu), nil
	}
}

func (s sysTeamMenuDo) Last() (*bizmodel.SysTeamMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamMenu), nil
	}
}

func (s sysTeamMenuDo) Find() ([]*bizmodel.SysTeamMenu, error) {
	result, err := s.DO.Find()
	return result.([]*bizmodel.SysTeamMenu), err
}

func (s sysTeamMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.SysTeamMenu, err error) {
	buf := make([]*bizmodel.SysTeamMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTeamMenuDo) FindInBatches(result *[]*bizmodel.SysTeamMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTeamMenuDo) Attrs(attrs ...field.AssignExpr) ISysTeamMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTeamMenuDo) Assign(attrs ...field.AssignExpr) ISysTeamMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTeamMenuDo) Joins(fields ...field.RelationField) ISysTeamMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTeamMenuDo) Preload(fields ...field.RelationField) ISysTeamMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTeamMenuDo) FirstOrInit() (*bizmodel.SysTeamMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamMenu), nil
	}
}

func (s sysTeamMenuDo) FirstOrCreate() (*bizmodel.SysTeamMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamMenu), nil
	}
}

func (s sysTeamMenuDo) FindByPage(offset int, limit int) (result []*bizmodel.SysTeamMenu, count int64, err error) {
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

func (s sysTeamMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTeamMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTeamMenuDo) Delete(models ...*bizmodel.SysTeamMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTeamMenuDo) withDO(do gen.Dao) *sysTeamMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
