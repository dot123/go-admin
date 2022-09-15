package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysDept struct {
	DeptID    uint64     `gorm:"column:dept_id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	ParentID  uint64     `gorm:"column:parent_id;type:bigint(20)"`
	DeptPath  string     `gorm:"column:dept_path;type:varchar(255)"`
	DeptName  string     `gorm:"column:dept_name;type:varchar(128)"`
	Sort      int        `gorm:"column:sort;type:tinyint(4)"`
	Leader    string     `gorm:"column:leader;type:varchar(128)"`
	Phone     string     `gorm:"column:phone;type:varchar(11)"`
	Email     string     `gorm:"column:email;type:varchar(64)"`
	Status    int        `gorm:"column:status;type:tinyint(1);default:0"`
	CreateBy  uint64     `gorm:"column:create_by;type:bigint(20)"`   // 创建者
	UpdateBy  uint64     `gorm:"column:update_by;type:bigint(20)"`   // 更新者
	CreatedAt types.Time `gorm:"column:created_at;type:datetime(3)"` // 创建时间
	UpdatedAt types.Time `gorm:"column:updated_at;type:datetime(3)"` // 最后更新时间
	DeletedAt types.Time `gorm:"column:deleted_at;type:datetime(3)"` // 删除时间
}

func GetSysDeptDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysDept))
}
