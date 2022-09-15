package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysApi struct {
	ID        uint64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 主键编码
	Handle    string     `gorm:"column:handle;type:varchar(128)"`                      // handle
	Title     string     `gorm:"column:title;type:varchar(128)"`                       // 标题
	Path      string     `gorm:"column:path;type:varchar(128)"`                        // 地址
	Type      string     `gorm:"column:type;type:varchar(16)"`                         // 接口类型
	Action    string     `gorm:"column:action;type:varchar(16)"`                       // 请求类型
	CreatedAt types.Time `gorm:"column:created_at;type:datetime(3)"`                   // 创建时间
	UpdatedAt types.Time `gorm:"column:updated_at;type:datetime(3)"`                   // 最后更新时间
	DeletedAt types.Time `gorm:"column:deleted_at;type:datetime(3)"`                   // 删除时间
	CreateBy  uint64     `gorm:"column:create_by;type:bigint(20)"`                     // 创建者
	UpdateBy  uint64     `gorm:"column:update_by;type:bigint(20)"`                     // 更新者
}

func GetSysApiDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysApi))
}
