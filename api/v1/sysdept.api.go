package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SysDeptApiSet = wire.NewSet(wire.Struct(new(SysDeptApi), "*"))

type SysDeptApi struct {
	SysDeptSrv *service.SysDeptSrv
}

// RegisterRoute 注册路由
func (a *SysDeptApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("dept", a.GetPage)
	r.GET("dept/:deptId", a.Get)
	r.POST("dept", a.Insert)
	r.PUT("dept", a.Update)
	r.DELETE("dept", a.Delete)
	r.GET("/deptTree", a.Get2Tree)
	r.GET("/roleDeptTreeselect/:roleId", a.GetDeptTreeRoleSelect)
}

// @Tags     部门
// @Summary  分页部门列表数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysDeptGetPageReq                   true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=[]schema.DeptLabel} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                        "失败结果"
// @Router   /dept [get]
func (a *SysDeptApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysDeptGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysDeptSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     部门
// @Summary  获取部门数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    deptId path     uint64                                 false "deptId"
// @Success  200    {object} ginx.ResponseData{data=schema.SysDept} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /dept/{deptId} [get]
func (a *SysDeptApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "deptId")

	result, err := a.SysDeptSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     部门
// @Summary  添加部门
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysDeptInsertReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /dept [post]
func (a *SysDeptApi) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysDeptInsertReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.CreateBy = contextx.FromUserID(ctx)

	err := a.SysDeptSrv.Insert(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     部门
// @Summary  修改部门
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysDeptUpdateReq true "body"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /dept [put]
func (a *SysDeptApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysDeptUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysDeptSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     部门
// @Summary  删除部门
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysDeptDeleteReq true "body"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /dept [delete]
func (a *SysDeptApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysDeptDeleteReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	p := contextx.FromDataPermission(ctx)
	req.RoleID = p.RoleID
	req.UserID = p.UserID

	err := a.SysDeptSrv.Delete(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     部门
// @Summary  左侧部门树
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysDeptGetPageReq                   true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=[]schema.DeptLabel} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                        "失败结果"
// @Router   /deptTree [get]
func (a *SysDeptApi) Get2Tree(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysDeptGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	list, err := a.SysDeptSrv.GetDeptTree(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, list)
}

// @Tags     部门
// @Summary  部门树形
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    roleId path     uint64                                                   false "roleId"
// @Success  200    {object} ginx.ResponseData{data=schema.GetDeptTreeRoleSelectResp} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                                      "失败结果"
// @Router   /roleDeptTreeselect/{roleId} [get]
func (a *SysDeptApi) GetDeptTreeRoleSelect(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.SysDeptSrv.GetDeptLabel(ctx)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	id := ginx.ParseParamID(c, "roleId")

	var menuIds *[]uint64
	if id != 0 {
		menuIds, err = a.SysDeptSrv.GetWithRoleId(ctx, id)
		if err != nil {
			ginx.ResError(c, err)
			return
		}
	}

	ginx.ResData(c, &schema.GetDeptTreeRoleSelectResp{
		Depts:       result,
		CheckedKeys: menuIds,
	})
}
