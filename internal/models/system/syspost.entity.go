package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysPost struct {
	PostID    uint64     `gorm:"column:post_id;type:bigint(20);primary_key;AUTO_INCREMENT"` //
	PostName  string     `gorm:"column:post_name;type:varchar(128)"`                        //
	PostCode  string     `gorm:"column:post_code;type:varchar(128)"`                        //
	Sort      int        `gorm:"column:sort;type:tinyint(4)"`                               //
	Status    int        `gorm:"column:status;type:tinyint(1);default:0"`                   // 状态 1停用 2启用
	Remark    string     `gorm:"column:remark;type:varchar(255)"`                           //
	CreateBy  uint64     `gorm:"column:create_by;type:bigint(20)"`                          // 创建者
	UpdateBy  uint64     `gorm:"column:update_by;type:bigint(20)"`                          // 更新者
	CreatedAt types.Time `gorm:"column:created_at;type:datetime(3)"`                        // 创建时间
	UpdatedAt types.Time `gorm:"column:updated_at;type:datetime(3)"`                        // 最后更新时间
	DeletedAt types.Time `gorm:"column:deleted_at;type:datetime(3)"`                        // 删除时间
}

func GetSysPostDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysPost))
}
