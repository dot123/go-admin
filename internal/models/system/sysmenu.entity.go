package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysMenu struct {
	MenuID     uint64     `gorm:"column:menu_id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	MenuName   string     `gorm:"column:menu_name;type:varchar(128)"`
	Title      string     `gorm:"column:title;type:varchar(128)"`
	Icon       string     `gorm:"column:icon;type:varchar(128)"`
	Path       string     `gorm:"column:path;type:varchar(128)"`
	Paths      string     `gorm:"column:paths;type:varchar(128)"`
	MenuType   string     `gorm:"column:menu_type;type:varchar(1)"`
	Action     string     `gorm:"column:action;type:varchar(16)"`
	Permission string     `gorm:"column:permission;type:varchar(255)"`
	ParentID   uint64     `gorm:"column:parent_id;type:bigint(20)"`
	NoCache    int        `gorm:"column:no_cache;type:tinyint(1)"`             // 1不缓存 2可缓存
	Breadcrumb int        `gorm:"column:breadcrumb;type:tinyint(1);default:0"` // 是否面包屑
	Component  string     `gorm:"column:component;type:varchar(255)"`          //
	Sort       int        `gorm:"column:sort;type:tinyint(4)"`                 //
	Visible    int        `gorm:"column:visible;type:tinyint(1);default:0"`    // 1不可见 2可见
	IsFrame    int        `gorm:"column:is_frame;type:tinyint(1);default:0"`   // 1不可见 2可见
	CreateBy   uint64     `gorm:"column:create_by;type:bigint(20)"`            // 创建者
	UpdateBy   uint64     `gorm:"column:update_by;type:bigint(20)"`            // 更新者
	CreatedAt  types.Time `gorm:"column:created_at;type:datetime(3)"`          // 创建时间
	UpdatedAt  types.Time `gorm:"column:updated_at;type:datetime(3)"`          // 最后更新时间
	DeletedAt  types.Time `gorm:"column:deleted_at;type:datetime(3)"`          // 删除时间
	SysApi     []*SysApi  `gorm:"many2many:sys_menu_api_rule"`
}

func GetSysMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysMenu))
}

type SysMenuSlice []*SysMenu

func (x SysMenuSlice) Len() int           { return len(x) }
func (x SysMenuSlice) Less(i, j int) bool { return x[i].Sort < x[j].Sort }
func (x SysMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
