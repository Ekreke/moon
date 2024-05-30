package bizmodel

import (
	"context"
	"encoding/json"

	"github.com/aide-cloud/moon/pkg/types"
	"github.com/aide-cloud/moon/pkg/vobj"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const TableNameDatasourceMetric = "datasource_metrics"

// DatasourceMetric mapped from table <datasource_metrics>
type DatasourceMetric struct {
	ID           uint32          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name         string          `gorm:"column:name;type:varchar(64);not null;comment:数据源名称" json:"name"`                                   // 数据源名称
	Category     vobj.MetricType `gorm:"column:category;type:int;not null;comment:指标类型（对应prometheus四种数据类型）" json:"category"`                // 指标类型（对应prometheus四种数据类型）
	Unit         string          `gorm:"column:unit;type:varchar(16);not null;comment:单位" json:"unit"`                                      // 单位
	Remark       string          `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`                                 // 备注
	DatasourceID uint32          `gorm:"column:datasource_id;type:int unsigned;not null;comment:所属数据源" json:"datasource_id"`                // 所属数据源
	CreatedAt    types.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt    types.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt    int64           `gorm:"column:deleted_at;type:bigint;not null;comment:删除时间" json:"deleted_at"`                             // 删除时间
	Labels       []*MetricLabel  `gorm:"foreignKey:MetricID" json:"labels"`
}

// String json string
func (c *DatasourceMetric) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *DatasourceMetric) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *DatasourceMetric) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// Create func
func (c *DatasourceMetric) Create(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Create(c).Error
}

// Update func
func (c *DatasourceMetric) Update(ctx context.Context, tx *gorm.DB, conds []gen.Condition) error {
	return tx.WithContext(ctx).Model(c).Where(conds).Updates(c).Error
}

// Delete func
func (c *DatasourceMetric) Delete(ctx context.Context, tx *gorm.DB, conds []gen.Condition) error {
	return tx.WithContext(ctx).Where(conds).Delete(c).Error
}

// TableName DatasourceMetric's table name
func (*DatasourceMetric) TableName() string {
	return TableNameDatasourceMetric
}