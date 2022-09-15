package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type SysUser struct {
	UserID    uint64     `gorm:"column:user_id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 编码
	Username  string     `gorm:"column:username;type:varchar(64)"`                          // 用户名
	Password  string     `gorm:"column:password;type:varchar(128)"`                         // 密码
	NickName  string     `gorm:"column:nick_name;type:varchar(128)"`                        // 昵称
	Phone     string     `gorm:"column:phone;type:varchar(11)"`                             // 手机号
	RoleID    uint64     `gorm:"column:role_id;type:bigint(20)"`                            // 角色ID
	Salt      string     `gorm:"column:salt;type:varchar(255)"`                             // 加盐
	Avatar    string     `gorm:"column:avatar;type:varchar(255)"`                           // 头像
	Sex       int        `gorm:"column:sex;type:tinyint(1)"`                                // 性别 1男 2女
	Email     string     `gorm:"column:email;type:varchar(128)"`                            // 邮箱
	DeptID    uint64     `gorm:"column:dept_id;type:bigint(20)"`                            // 部门
	PostID    uint64     `gorm:"column:post_id;type:bigint(20)"`                            // 岗位
	Remark    string     `gorm:"column:remark;type:varchar(255)"`                           // 备注
	Status    int        `gorm:"column:status;type:tinyint(1);default:0"`                   // 状态 1停用 2启用
	CreateBy  uint64     `gorm:"column:create_by;type:bigint(20)"`                          // 创建者
	UpdateBy  uint64     `gorm:"column:update_by;type:bigint(20)"`                          // 更新者
	CreatedAt types.Time `gorm:"column:created_at;type:datetime(3)"`                        // 创建时间
	UpdatedAt types.Time `gorm:"column:updated_at;type:datetime(3)"`                        // 最后更新时间
	DeletedAt types.Time `gorm:"column:deleted_at;type:datetime(3)"`                        // 删除时间
	DeptIds   []uint64   `gorm:"-"`
	PostIds   []uint64   `gorm:"-"`
	RoleIds   []uint64   `gorm:"-"`
	Dept      *SysDept
}

func GetSysUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(SysUser))
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.DeptIds = []uint64{e.DeptID}
	e.PostIds = []uint64{e.PostID}
	e.RoleIds = []uint64{e.RoleID}
	return nil
}
