package schema

type SysPost struct {
	PostID    uint64 `json:"postId"`   //岗位编号
	PostName  string `json:"postName"` //岗位名称
	PostCode  string `json:"postCode"` //岗位代码
	Sort      int    `json:"sort"`     //岗位排序
	Status    int    `json:"status"`   //状态
	Remark    string `json:"remark"`   //描述
	CreateBy  uint64 `json:"createBy"` // 创建者
	UpdateBy  uint64 `json:"updateBy"` // 更新者
	DataScope string `json:"dataScope"`
	Params    string `json:"params"`
}

type SysPostPageReq struct {
	Pagination
	PostID   uint64 `form:"postID"`   // id
	PostName string `form:"postName"` // 名称
	PostCode string `form:"postCode"` // 编码
	Sort     int    `form:"sort"`     // 排序
	Status   int    `form:"status"`   // 状态
	Remark   string `form:"remark"`   // 备注
}

type SysPostPageResp struct {
	Data  []*SysPost
	Total int64
}

type SysPostInsertReq struct {
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
	CreateBy uint64 `json:"createBy"` // 创建者
}

type SysPostUpdateReq struct {
	PostID   uint64 `json:"id"`
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int    `json:"sort"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
	UpdateBy uint64 `json:"updateBy"` // 更新者
}

type SysPostDeleteReq struct {
	Ids []uint64 `json:"ids"`
}
