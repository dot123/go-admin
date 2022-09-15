package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SysRoleApiSet = wire.NewSet(wire.Struct(new(SysRoleApi), "*"))

type SysRoleApi struct {
	SysRoleSrv *service.SysRoleSrv
}

// RegisterRoute 注册路由
func (a *SysRoleApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("role", a.GetPage)
	r.GET("role/:roleId", a.Get)
	r.POST("role", a.Insert)
	r.PUT("role", a.Update)
	r.DELETE("role/:roleId", a.Delete)
	r.PUT("role-status", a.Update2Status)
	r.PUT("role-datascope", a.Update2DataScope)
}

// @Tags     角色/Role
// @Summary  角色列表数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysRoleGetPageReq                          true "body"
// @Success  200  {object} ginx.ResponseData{data=schema.SysRoleGetPageResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                               "失败结果"
// @Router   /role [get]
func (a *SysRoleApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysRoleGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysRoleSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     角色/Role
// @Summary  获取Role数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    roleId path     uint64                                 false "roleId"
// @Success  200    {object} ginx.ResponseData{data=schema.SysRole} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /role/{roleId} [get]
func (a *SysRoleApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "roleId")
	result, err := a.SysRoleSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     角色/Role
// @Summary  创建角色
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysRoleInsertReq true "data"
// @Success  200    {object} ginx.ResponseData{} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{} "失败结果"
// @Router   /role [post]
func (a *SysRoleApi) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysRoleInsertReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.CreateBy = contextx.FromUserID(ctx)
	if req.Status == 0 {
		req.Status = 2
	}

	err := a.SysRoleSrv.Insert(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     角色/Role
// @Summary  修改用户角色
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysRoleUpdateReq true "data"
// @Success  200  {object} ginx.ResponseData{}    "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}    "失败结果"
// @Router   /role [put]
func (a *SysRoleApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysRoleUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)
	if req.Status == 0 {
		req.Status = 2
	}

	err := a.SysRoleSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     角色/Role
// @Summary  删除用户角色
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    roleId path     uint64              false "roleId"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /role/{roleId} [delete]
func (a *SysRoleApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "roleId")
	err := a.SysRoleSrv.Delete(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     角色/Role
// @Summary  修改用户角色状态
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.UpdateStatusReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /role-status [put]
func (a *SysRoleApi) Update2Status(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.UpdateStatusReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)
	if req.Status == 0 {
		req.Status = 2
	}

	err := a.SysRoleSrv.UpdateStatus(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     角色/Role
// @Summary  更新角色数据权限
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.RoleDataScopeReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /role-datascope [put]
func (a *SysRoleApi) Update2DataScope(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.RoleDataScopeReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysRoleSrv.UpdateDataScope(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
