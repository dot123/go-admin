package schema

import (
	"github.com/dot123/gin-gorm-admin/internal/validate"
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type SysDept struct {
	DeptID    uint64     `json:"deptId"`    //部门编码
	ParentID  uint64     `json:"parentId"`  //上级部门
	DeptPath  string     `json:"deptPath"`  //
	DeptName  string     `json:"deptName"`  //部门名称
	Sort      int        `json:"sort"`      //排序
	Leader    string     `json:"leader"`    //负责人
	Phone     string     `json:"phone"`     //手机
	Email     string     `json:"email"`     //邮箱
	Status    int        `json:"status"`    //状态
	CreatedAt types.Time `json:"createdAt"` // 创建时间
	UpdatedAt types.Time `json:"updatedAt"` // 最后更新时间
	CreateBy  uint64     `json:"createBy"`  // 创建者
	UpdateBy  uint64     `json:"updateBy"`  // 更新者
	DataScope string     `json:"dataScope"`
	Params    string     `json:"params"`
	Children  []*SysDept `json:"children"`
}

type SysDeptGetPageReq struct {
	DeptID   uint64 `form:"deptID"`   //id
	ParentID uint64 `form:"parentID"` //上级部门
	DeptPath string `form:"deptPath"` //路径
	DeptName string `form:"deptName"` //部门名称
	Sort     int    `form:"sort"`     //排序
	Leader   string `form:"leader"`   //负责人
	Phone    string `form:"phone"`    //手机
	Email    string `form:"email"`    //邮箱
	Status   int    `form:"status"`   //状态
}

type SysDeptInsertReq struct {
	ParentID uint64 `json:"parentId" validate:"required"`           // 上级部门
	DeptPath string `json:"deptPath"`                               // 路径
	DeptName string `json:"deptName" validate:"required"`           // 部门名称
	Sort     int    `json:"sort" validate:"required"`               // 排序
	Leader   string `json:"leader" validate:"required"`             // 负责人
	Phone    string `json:"phone" validate:"required"`              // 手机
	Email    string `json:"email" validate:"required"`              // 邮箱
	Status   int    `json:"status" validate:"required" default:"2"` // 状态
	CreateBy uint64 `json:"createBy"`                               // 创建者
}

func (m *SysDeptInsertReq) Verify() string {
	messages := map[string]string{
		"ParentID.required": "上级部门不能为空",
		"DeptName.required": "部门名称不能为空",
		"Sort.required":     "排序不能为空",
		"Leader.required":   "负责人不能为空",
		"Phone.required":    "手机号码不能为空",
		"Email.required":    "邮箱不能为空",
		"Status.required":   "状态不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type SysDeptUpdateReq struct {
	DeptID   uint64 `json:"deptId"`                                 // 编码
	ParentID uint64 `json:"parentId" validate:"required"`           // 上级部门
	DeptPath string `json:"deptPath"`                               // 路径
	DeptName string `json:"deptName" validate:"required"`           // 部门名称
	Sort     int    `json:"sort" validate:"required"`               // 排序
	Leader   string `json:"leader" validate:"required"`             // 负责人
	Phone    string `json:"phone" validate:"required"`              // 手机
	Email    string `json:"email" validate:"required"`              // 邮箱
	Status   int    `json:"status" validate:"required" default:"2"` // 状态
	UpdateBy uint64 `json:"updateBy"`                               // 更新者
}

func (m *SysDeptUpdateReq) Verify() string {
	messages := map[string]string{
		"ParentID.required": "上级部门不能为空",
		"DeptName.required": "部门名称不能为空",
		"Sort.required":     "排序不能为空",
		"Leader.required":   "负责人不能为空",
		"Phone.required":    "手机号码不能为空",
		"Email.required":    "邮箱不能为空",
		"Status.required":   "状态不能为空",
	}
	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type SysDeptDeleteReq struct {
	Ids    []uint64 `json:"ids"`
	RoleID uint64   `json:"-"`
	UserID uint64   `json:"-"`
}

type DeptLabel struct {
	ID       uint64       `json:"id"`
	Label    string       `json:"label"`
	Children []*DeptLabel `json:"children"`
}

type DeptIDList struct {
	DeptID uint64 `json:"DeptId"`
}

type GetDeptTreeRoleSelectResp struct {
	Depts       *[]*DeptLabel `json:"depts"`
	CheckedKeys *[]uint64     `json:"checkedKeys"`
}
