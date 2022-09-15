package schema

import (
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type SysMenu struct {
	MenuID     uint64     `json:"menuId"`
	MenuName   string     `json:"menuName"`
	Title      string     `json:"title"`
	Icon       string     `json:"icon"`
	Path       string     `json:"path"`
	Paths      string     `json:"paths"`
	MenuType   string     `json:"menuType"`
	Action     string     `json:"action"`
	Permission string     `json:"permission"`
	ParentID   uint64     `json:"parentId"`
	NoCache    int        `json:"noCache"`
	Breadcrumb int        `json:"breadcrumb"`
	Component  string     `json:"component"`
	Sort       int        `json:"sort"`
	Visible    int        `json:"visible"`
	IsFrame    int        `json:"isFrame"`
	CreatedAt  types.Time `json:"createdAt"` // 创建时间
	UpdatedAt  types.Time `json:"updatedAt"` // 最后更新时间
	CreateBy   uint64     `json:"createBy"`  // 创建者
	UpdateBy   uint64     `json:"updateBy"`  // 更新者
	SysApi     []*SysApi  `json:"sysApi"`
	Apis       []uint64   `json:"apis"`
	DataScope  string     `json:"dataScope"`
	Params     string     `json:"params"`
	RoleID     uint64     `json:"roleId"`
	Children   []*SysMenu `json:"children,omitempty"`
	IsSelect   bool       `json:"is_select"`
}

type SysMenuGetPageReq struct {
	Pagination
	Title   string `form:"title"`   // 菜单名称
	Visible int    `form:"visible"` // 显示状态
}

// Menu 菜单中的类型枚举值
const (
	// Directory 目录
	Directory string = "M"
	// Menu 菜单
	Menu string = "C"
	// Button 按钮
	Button string = "F"
)

type SysMenuInsertReq struct {
	MenuName   string `json:"menuName"`   //菜单name
	Title      string `json:"title"`      //显示名称
	Icon       string `json:"icon"`       //图标
	Path       string `json:"path"`       //路径
	Paths      string `json:"paths"`      //id路径
	MenuType   string `json:"menuType"`   //菜单类型
	Action     string `json:"action"`     //请求方式
	Permission string `json:"permission"` //权限编码
	ParentID   uint64 `json:"parentId"`   //上级菜单
	NoCache    int    `json:"noCache"`    //是否缓存
	Breadcrumb int    `json:"breadcrumb"` //是否面包屑
	Component  string `json:"component"`  //组件
	Sort       int    `json:"sort"`       //排序
	Visible    int    `json:"visible"`    //是否显示
	IsFrame    int    `json:"isFrame"`    //是否frame
	CreateBy   uint64 `json:"createBy"`
}

type SysMenuUpdateReq struct {
	MenuID     uint64   `json:"id"`       // 编码
	MenuName   string   `json:"menuName"` //菜单name
	Title      string   `json:"title"`    //显示名称
	Icon       string   `json:"icon"`     //图标
	Path       string   `json:"path"`     //路径
	Paths      string   `json:"paths"`    //id路径
	MenuType   string   `json:"menuType"` //菜单类型
	Apis       []uint64 `json:"apis"`
	Action     string   `json:"action"`     //请求方式
	Permission string   `json:"permission"` //权限编码
	ParentID   uint64   `json:"parentId"`   //上级菜单
	NoCache    int      `json:"noCache"`    //是否缓存
	Breadcrumb int      `json:"breadcrumb"` //是否面包屑
	Component  string   `json:"component"`  //组件
	Sort       int      `json:"sort"`       //排序
	Visible    int      `json:"visible"`    //是否显示
	IsFrame    string   `json:"isFrame"`    //是否frame
	UpdateBy   uint64   `json:"updateBy"`
}

type SysMenuDeleteReq struct {
	Ids    []uint64 `json:"ids"`
	RoleID uint64   `json:"-"`
}

type MenuLabel struct {
	ID       uint64       `json:"id,omitempty"`
	Label    string       `json:"label,omitempty"`
	Children []*MenuLabel `json:"children,omitempty"`
}

type GetMenuTreeSelectResp struct {
	Menus       *[]*MenuLabel `json:"menus"`
	CheckedKeys *[]uint64     `json:"checkedKeys"`
}
