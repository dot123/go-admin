package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysRole struct {
	RoleID    uint64      `gorm:"column:role_id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 编码
	RoleName  string      `gorm:"column:role_name;type:varchar(128)"`                        //
	Status    int         `gorm:"column:status;type:tinyint(1);default:0"`                   // 状态 1停用 2启用
	RoleKey   string      `gorm:"column:role_key;type:varchar(128)"`                         //
	RoleSort  int         `gorm:"column:role_sort;type:tinyint(4)"`                          //
	Flag      string      `gorm:"column:flag;type:varchar(128)"`                             //
	Remark    string      `gorm:"column:remark;type:varchar(255)"`                           //
	Admin     int         `gorm:"column:admin;type:tinyint(1)"`                              // 1不是管理员 2是管理员
	DataScope string      `gorm:"column:data_scope;type:varchar(128)"`                       // 1全部权限 2自定义数据权限 3本部门数据权限 4本部门及以下数据权限 5仅本人数据权限
	CreateBy  uint64      `gorm:"column:create_by;type:bigint(20)"`                          // 创建者
	UpdateBy  uint64      `gorm:"column:update_by;type:bigint(20)"`                          // 更新者
	CreatedAt types.Time  `gorm:"column:created_at;type:datetime(3)"`                        // 创建时间
	UpdatedAt types.Time  `gorm:"column:updated_at;type:datetime(3)"`                        // 最后更新时间
	DeletedAt types.Time  `gorm:"column:deleted_at;type:datetime(3)"`                        // 删除时间
	SysDept   *[]*SysDept `gorm:"many2many:sys_role_dept;foreignKey:RoleID;joinForeignKey:role_id;references:DeptID;joinReferences:dept_id;"`
	SysMenu   *[]*SysMenu `gorm:"many2many:sys_role_menu;foreignKey:RoleID;joinForeignKey:role_id;references:MenuID;joinReferences:menu_id;"`
}

func GetSysRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysRole))
}
