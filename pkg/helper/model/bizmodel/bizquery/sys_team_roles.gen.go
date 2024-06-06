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

func newSysTeamRole(db *gorm.DB, opts ...gen.DOOption) sysTeamRole {
	_sysTeamRole := sysTeamRole{}

	_sysTeamRole.sysTeamRoleDo.UseDB(db, opts...)
	_sysTeamRole.sysTeamRoleDo.UseModel(&bizmodel.SysTeamRole{})

	tableName := _sysTeamRole.sysTeamRoleDo.TableName()
	_sysTeamRole.ALL = field.NewAsterisk(tableName)
	_sysTeamRole.ID = field.NewUint32(tableName, "id")
	_sysTeamRole.CreatedAt = field.NewField(tableName, "created_at")
	_sysTeamRole.UpdatedAt = field.NewField(tableName, "updated_at")
	_sysTeamRole.DeletedAt = field.NewUint(tableName, "deleted_at")
	_sysTeamRole.TeamID = field.NewUint32(tableName, "team_id")
	_sysTeamRole.Name = field.NewString(tableName, "name")
	_sysTeamRole.Status = field.NewInt(tableName, "status")
	_sysTeamRole.Remark = field.NewString(tableName, "remark")
	_sysTeamRole.Apis = sysTeamRoleManyToManyApis{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Apis", "bizmodel.SysTeamAPI"),
	}

	_sysTeamRole.fillFieldMap()

	return _sysTeamRole
}

type sysTeamRole struct {
	sysTeamRoleDo

	ALL       field.Asterisk
	ID        field.Uint32
	CreatedAt field.Field
	UpdatedAt field.Field
	DeletedAt field.Uint
	TeamID    field.Uint32
	Name      field.String
	Status    field.Int
	Remark    field.String
	Apis      sysTeamRoleManyToManyApis

	fieldMap map[string]field.Expr
}

func (s sysTeamRole) Table(newTableName string) *sysTeamRole {
	s.sysTeamRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTeamRole) As(alias string) *sysTeamRole {
	s.sysTeamRoleDo.DO = *(s.sysTeamRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTeamRole) updateTableName(table string) *sysTeamRole {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "id")
	s.CreatedAt = field.NewField(table, "created_at")
	s.UpdatedAt = field.NewField(table, "updated_at")
	s.DeletedAt = field.NewUint(table, "deleted_at")
	s.TeamID = field.NewUint32(table, "team_id")
	s.Name = field.NewString(table, "name")
	s.Status = field.NewInt(table, "status")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysTeamRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTeamRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["team_id"] = s.TeamID
	s.fieldMap["name"] = s.Name
	s.fieldMap["status"] = s.Status
	s.fieldMap["remark"] = s.Remark

}

func (s sysTeamRole) clone(db *gorm.DB) sysTeamRole {
	s.sysTeamRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTeamRole) replaceDB(db *gorm.DB) sysTeamRole {
	s.sysTeamRoleDo.ReplaceDB(db)
	return s
}

type sysTeamRoleManyToManyApis struct {
	db *gorm.DB

	field.RelationField
}

func (a sysTeamRoleManyToManyApis) Where(conds ...field.Expr) *sysTeamRoleManyToManyApis {
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

func (a sysTeamRoleManyToManyApis) WithContext(ctx context.Context) *sysTeamRoleManyToManyApis {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sysTeamRoleManyToManyApis) Session(session *gorm.Session) *sysTeamRoleManyToManyApis {
	a.db = a.db.Session(session)
	return &a
}

func (a sysTeamRoleManyToManyApis) Model(m *bizmodel.SysTeamRole) *sysTeamRoleManyToManyApisTx {
	return &sysTeamRoleManyToManyApisTx{a.db.Model(m).Association(a.Name())}
}

type sysTeamRoleManyToManyApisTx struct{ tx *gorm.Association }

func (a sysTeamRoleManyToManyApisTx) Find() (result []*bizmodel.SysTeamAPI, err error) {
	return result, a.tx.Find(&result)
}

func (a sysTeamRoleManyToManyApisTx) Append(values ...*bizmodel.SysTeamAPI) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sysTeamRoleManyToManyApisTx) Replace(values ...*bizmodel.SysTeamAPI) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sysTeamRoleManyToManyApisTx) Delete(values ...*bizmodel.SysTeamAPI) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sysTeamRoleManyToManyApisTx) Clear() error {
	return a.tx.Clear()
}

func (a sysTeamRoleManyToManyApisTx) Count() int64 {
	return a.tx.Count()
}

type sysTeamRoleDo struct{ gen.DO }

type ISysTeamRoleDo interface {
	gen.SubQuery
	Debug() ISysTeamRoleDo
	WithContext(ctx context.Context) ISysTeamRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTeamRoleDo
	WriteDB() ISysTeamRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTeamRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTeamRoleDo
	Not(conds ...gen.Condition) ISysTeamRoleDo
	Or(conds ...gen.Condition) ISysTeamRoleDo
	Select(conds ...field.Expr) ISysTeamRoleDo
	Where(conds ...gen.Condition) ISysTeamRoleDo
	Order(conds ...field.Expr) ISysTeamRoleDo
	Distinct(cols ...field.Expr) ISysTeamRoleDo
	Omit(cols ...field.Expr) ISysTeamRoleDo
	Join(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo
	Group(cols ...field.Expr) ISysTeamRoleDo
	Having(conds ...gen.Condition) ISysTeamRoleDo
	Limit(limit int) ISysTeamRoleDo
	Offset(offset int) ISysTeamRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamRoleDo
	Unscoped() ISysTeamRoleDo
	Create(values ...*bizmodel.SysTeamRole) error
	CreateInBatches(values []*bizmodel.SysTeamRole, batchSize int) error
	Save(values ...*bizmodel.SysTeamRole) error
	First() (*bizmodel.SysTeamRole, error)
	Take() (*bizmodel.SysTeamRole, error)
	Last() (*bizmodel.SysTeamRole, error)
	Find() ([]*bizmodel.SysTeamRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.SysTeamRole, err error)
	FindInBatches(result *[]*bizmodel.SysTeamRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*bizmodel.SysTeamRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTeamRoleDo
	Assign(attrs ...field.AssignExpr) ISysTeamRoleDo
	Joins(fields ...field.RelationField) ISysTeamRoleDo
	Preload(fields ...field.RelationField) ISysTeamRoleDo
	FirstOrInit() (*bizmodel.SysTeamRole, error)
	FirstOrCreate() (*bizmodel.SysTeamRole, error)
	FindByPage(offset int, limit int) (result []*bizmodel.SysTeamRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTeamRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTeamRoleDo) Debug() ISysTeamRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTeamRoleDo) WithContext(ctx context.Context) ISysTeamRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTeamRoleDo) ReadDB() ISysTeamRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTeamRoleDo) WriteDB() ISysTeamRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTeamRoleDo) Session(config *gorm.Session) ISysTeamRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTeamRoleDo) Clauses(conds ...clause.Expression) ISysTeamRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTeamRoleDo) Returning(value interface{}, columns ...string) ISysTeamRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTeamRoleDo) Not(conds ...gen.Condition) ISysTeamRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTeamRoleDo) Or(conds ...gen.Condition) ISysTeamRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTeamRoleDo) Select(conds ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTeamRoleDo) Where(conds ...gen.Condition) ISysTeamRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTeamRoleDo) Order(conds ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTeamRoleDo) Distinct(cols ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTeamRoleDo) Omit(cols ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTeamRoleDo) Join(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTeamRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTeamRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTeamRoleDo) Group(cols ...field.Expr) ISysTeamRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTeamRoleDo) Having(conds ...gen.Condition) ISysTeamRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTeamRoleDo) Limit(limit int) ISysTeamRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTeamRoleDo) Offset(offset int) ISysTeamRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTeamRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTeamRoleDo) Unscoped() ISysTeamRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTeamRoleDo) Create(values ...*bizmodel.SysTeamRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTeamRoleDo) CreateInBatches(values []*bizmodel.SysTeamRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTeamRoleDo) Save(values ...*bizmodel.SysTeamRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTeamRoleDo) First() (*bizmodel.SysTeamRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamRole), nil
	}
}

func (s sysTeamRoleDo) Take() (*bizmodel.SysTeamRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamRole), nil
	}
}

func (s sysTeamRoleDo) Last() (*bizmodel.SysTeamRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamRole), nil
	}
}

func (s sysTeamRoleDo) Find() ([]*bizmodel.SysTeamRole, error) {
	result, err := s.DO.Find()
	return result.([]*bizmodel.SysTeamRole), err
}

func (s sysTeamRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.SysTeamRole, err error) {
	buf := make([]*bizmodel.SysTeamRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTeamRoleDo) FindInBatches(result *[]*bizmodel.SysTeamRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTeamRoleDo) Attrs(attrs ...field.AssignExpr) ISysTeamRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTeamRoleDo) Assign(attrs ...field.AssignExpr) ISysTeamRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTeamRoleDo) Joins(fields ...field.RelationField) ISysTeamRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTeamRoleDo) Preload(fields ...field.RelationField) ISysTeamRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTeamRoleDo) FirstOrInit() (*bizmodel.SysTeamRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamRole), nil
	}
}

func (s sysTeamRoleDo) FirstOrCreate() (*bizmodel.SysTeamRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.SysTeamRole), nil
	}
}

func (s sysTeamRoleDo) FindByPage(offset int, limit int) (result []*bizmodel.SysTeamRole, count int64, err error) {
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

func (s sysTeamRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTeamRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTeamRoleDo) Delete(models ...*bizmodel.SysTeamRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTeamRoleDo) withDO(do gen.Dao) *sysTeamRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
