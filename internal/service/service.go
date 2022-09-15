package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAuthSrv,
	FileSet,
	MonitorSet,
	NewMsgSrv,
	SysApiSet,
	SysUserSet,
	SysRoleSet,
	SysPostSet,
	SysDeptSet,
	SysMenuSet,
)
