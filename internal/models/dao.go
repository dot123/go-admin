package models

import (
	"github.com/dot123/gin-gorm-admin/internal/models/file"
	"github.com/dot123/gin-gorm-admin/internal/models/msg"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepoSet = wire.NewSet(
	util.TransSet,
	file.FileSet,
	msg.MsgSet,
	system.SysApiSet,
	system.SysUserSet,
	system.SysRoleSet,
	system.SysPostSet,
	system.SysDeptSet,
	system.SysMenuSet,
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		new(file.File),
		new(msg.Notice),
		new(system.SysRole),
		new(system.SysUser),
		new(system.SysDept),
		new(system.SysMenu),
		new(system.SysPost),
		new(system.SysApi),
	)
	return err
}
