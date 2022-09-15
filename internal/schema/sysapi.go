package schema

import (
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type SysApi struct {
	ID        uint64     `json:"id"`        // 主键编码
	Handle    string     `json:"handle"`    // handle
	Title     string     `json:"title"`     // 标题
	Path      string     `json:"path"`      // 地址
	Type      string     `json:"type"`      // 接口类型
	Action    string     `json:"action"`    // 请求类型
	CreatedAt types.Time `json:"createdAt"` // 创建时间
	UpdatedAt types.Time `json:"updatedAt"` // 最后更新时间
	DeletedAt types.Time `json:"deletedAt"` // 删除时间
	CreateBy  uint64     `json:"createBy"`  // 创建者
	UpdateBy  uint64     `json:"updateBy"`  // 更新者
}

type SysApiGetPageReq struct {
	Pagination
	Title    string `form:"title"`
	Path     string `form:"path"`
	Action   string `form:"action"`
	ParentId string `form:"parentId"`
	SysApiOrder
}

type SysApiOrder struct {
	TitleOrder     string `form:"titleOrder"`
	PathOrder      string `form:"pathOrder"`
	CreatedAtOrder string `form:"createdAtOrder"`
}

type SysApiGetPageResp struct {
	Data  []*SysApi
	Total int64
}

type SysApiUpdateReq struct {
	ID       uint64 `json:"id"`
	Handle   string `json:"handle"`
	Title    string `json:"title"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Action   string `json:"action"`
	UpdateBy uint64 `json:"updateBy"`
}

type SysApiDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type Router struct {
	HttpMethod, RelativePath, Handler string
}
