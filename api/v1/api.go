package v1

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	AuthApiSet,
	FileApiSet,
	MonitorApiSet,
	MsgApiSet,
	SysApiSet,
	SysUserApiSet,
	SysRoleApiSet,
	SysPostApiSet,
	SysDeptApiSet,
	SysMenuApiSet,
)
