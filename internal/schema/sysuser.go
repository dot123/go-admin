package schema

import (
	"github.com/dot123/gin-gorm-admin/internal/validate"
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type SysUser struct {
	UserID    uint64     `json:"userId"`
	Username  string     `json:"username"`
	NickName  string     `json:"nickName"`
	Phone     string     `json:"phone"`
	RoleID    uint64     `json:"roleId"`
	Avatar    string     `json:"avatar"`
	Sex       int        `json:"sex"`
	Email     string     `json:"email"`
	DeptID    uint64     `json:"deptId"`
	PostID    uint64     `json:"postId"`
	Remark    string     `json:"remark"`
	Status    int        `json:"status"`
	DeptIds   []uint64   `json:"deptIds"`
	PostIds   []uint64   `json:"postIds"`
	RoleIds   []uint64   `json:"roleIds" `
	Dept      *SysDept   `json:"dept"`
	CreatedAt types.Time `json:"createdAt"` // 创建时间
	UpdatedAt types.Time `json:"updatedAt"` // 最后更新时间
	CreateBy  uint64     `json:"createBy"`  // 创建者
	UpdateBy  uint64     `json:"updateBy"`  // 更新者
}

type SysUserGetPageReq struct {
	Pagination
	UserID   uint64 `form:"userID"`
	Username string `form:"username"`
	NickName string `form:"nickName"`
	Phone    string `form:"phone"`
	RoleID   uint64 `form:"roleID"`
	Sex      int    `form:"sex"`
	Email    string `form:"email"`
	PostID   uint64 `form:"postID"`
	Status   int    `form:"status"`
	DeptJoin
	SysUserOrder
}

type SysUserOrder struct {
	UserIdOrder    string `form:"userIdOrder"`
	UsernameOrder  string `form:"usernameOrder"`
	StatusOrder    string `form:"statusOrder"`
	CreatedAtOrder string `form:"createdAtOrder"`
}

type DeptJoin struct {
	DeptId uint64 `form:"deptId"`
}

type SysUserGetPageResp struct {
	Data  []*SysUser
	Total int64
}

type SysUserInsertReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	RoleID   uint64 `json:"roleId"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	Email    string `json:"email" validate:"required"`
	DeptID   uint64 `json:"deptId" validate:"required"`
	PostID   uint64 `json:"postId"`
	Remark   string `json:"remark"`
	Status   int    `json:"status" validate:"required" default:"2"`
	CreateBy uint64 `json:"createBy"`
}

func (m *SysUserInsertReq) Verify() string {
	messages := map[string]string{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
		"NickName.required": "用户昵称不能为空",
		"Phone.required":    "手机号码不能为空",
		"Email.required":    "邮箱不能为空",
		"DeptID.required":   "归属部门不能为空",
		"Status.required":   "状态不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type SysUserUpdateReq struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"username" validate:"required"`
	NickName string `json:"nickName" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	RoleID   uint64 `json:"roleId"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	Email    string `json:"email" validate:"required"`
	DeptID   uint64 `json:"deptId" validate:"required"`
	PostID   uint64 `json:"postId"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
	UpdateBy uint64 `json:"updateBy"`
}

func (m *SysUserUpdateReq) Verify() string {
	messages := map[string]string{
		"Username.required": "用户名不能为空",
		"NickName.required": "用户昵称不能为空",
		"Phone.required":    "手机号码不能为空",
		"Email.required":    "邮箱不能为空",
		"DeptID.required":   "归属部门不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type SysUserDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type UpdateSysUserAvatarReq struct {
	UserID   uint64 `json:"userId"`
	Avatar   string `json:"avatar"`
	UpdateBy uint64 `json:"updateBy"`
}

type UpdateSysUserStatusReq struct {
	UserID   uint64 `json:"userId"`
	Status   int    `json:"status"`
	UpdateBy uint64 `json:"updateBy"`
}

type ResetSysUserPwdReq struct {
	UserID   uint64 `json:"userId"`
	Password string `json:"password"`
	UpdateBy uint64 `json:"updateBy"`
}

type PassWord struct {
	NewPassword string `json:"newPassword" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}

func (m *PassWord) Verify() string {
	messages := map[string]string{
		"NewPassword.required": "新密码不能为空",
		"OldPassword.required": "旧密码不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type SysUserById struct {
	UserID uint64 `json:"userId"`
}

type GetProfileResp struct {
	User  *SysUser    `json:"user"`
	Roles *[]*SysRole `json:"roles"`
	Posts *[]*SysPost `json:"posts"`
}
