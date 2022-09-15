package schema

import (
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type SysRole struct {
	RoleID    uint64      `json:"roleId"`   // 角色编码
	RoleName  string      `json:"roleName"` // 角色名称
	Status    int         `json:"status"`   //
	RoleKey   string      `json:"roleKey"`  //角色代码
	RoleSort  int         `json:"roleSort"` //角色排序
	Flag      string      `json:"flag"`     //
	Remark    string      `json:"remark"`   //备注
	Admin     int         `json:"admin"`
	DataScope string      `json:"dataScope"`
	Params    string      `json:"params"`
	MenuIds   *[]uint64   `json:"menuIds"`
	SysMenu   *[]*SysMenu `json:"sysMenu"`
	CreatedAt types.Time  `json:"createdAt"` // 创建时间
	UpdatedAt types.Time  `json:"updatedAt"` // 最后更新时间
	CreateBy  uint64      `json:"createBy"`  // 创建者
	UpdateBy  uint64      `json:"updateBy"`  // 更新者
}

type SysRoleGetPageReq struct {
	Pagination
	RoleID    uint64 `form:"roleID"`   // 角色编码
	RoleName  string `form:"roleName"` // 角色名称
	Status    int    `form:"status"`   // 状态
	RoleKey   string `form:"roleKey"`  // 角色代码
	RoleSort  int    `form:"roleSort"` // 角色排序
	Flag      string `form:"flag"`     // 标记
	Remark    string `form:"remark"`   // 备注
	Admin     int    `form:"admin"`    // 1不是管理员 2是管理员
	DataScope string `form:"dataScope"`
}

type SysRoleOrder struct {
	RoleIdOrder    string `form:"roleIdOrder"`
	RoleNameOrder  string `form:"roleNameOrder"`
	RoleSortOrder  string `form:"usernameOrder"`
	StatusOrder    string `form:"statusOrder"`
	CreatedAtOrder string `form:"createdAtOrder"`
}

type SysRoleGetPageResp struct {
	Data  []*SysRole
	Total int64
}

type SysRoleInsertReq struct {
	RoleName  string   `json:"roleName"` // 角色名称
	Status    int      `json:"status"`   // 状态
	RoleKey   string   `json:"roleKey"`  // 角色代码
	RoleSort  int      `json:"roleSort"` // 角色排序
	Flag      string   `json:"flag"`     // 标记
	Remark    string   `json:"remark"`   // 备注
	Admin     int      `json:"admin"`
	DataScope string   `json:"dataScope"`
	MenuIds   []uint64 `json:"menuIds"`
	CreateBy  uint64   `json:"createBy"` // 创建者
}

type SysRoleUpdateReq struct {
	RoleID    uint64   `json:"id"`       // 角色编码
	RoleName  string   `json:"roleName"` // 角色名称
	Status    int      `json:"status"`   // 状态
	RoleKey   string   `json:"roleKey"`  // 角色代码
	RoleSort  int      `json:"roleSort"` // 角色排序
	Flag      string   `json:"flag"`     // 标记
	Remark    string   `json:"remark"`   // 备注
	Admin     int      `json:"admin"`
	DataScope string   `json:"dataScope"`
	MenuIds   []uint64 `json:"menuIds"`
	UpdateBy  uint64   `json:"updateBy"` // 更新者
}

type RoleDataScopeReq struct {
	RoleID    uint64   `json:"id" validate:"required"` // 角色编码
	DataScope string   `json:"dataScope" validate:"required"`
	DeptIds   []uint64 `json:"deptIds"`
	UpdateBy  uint64   `json:"updateBy"` // 更新者
}

type UpdateStatusReq struct {
	RoleID   uint64 `json:"id"`       // 角色编码
	Status   int    `json:"status"`   // 状态
	UpdateBy uint64 `json:"updateBy"` // 更新者
}
