package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SysMenuApiSet = wire.NewSet(wire.Struct(new(SysMenuApi), "*"))

type SysMenuApi struct {
	SysMenuSrv *service.SysMenuSrv
	SysRoleSrv *service.SysRoleSrv
}

// RegisterRoute 注册路由
func (a *SysMenuApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("menu", a.GetPage)
	r.GET("menu/:id", a.Get)
	r.POST("menu", a.Insert)
	r.PUT("menu", a.Update)
	r.DELETE("menu", a.Delete)
	r.GET("menurole", a.GetMenuRole)
	r.GET("roleMenuTreeselect/:roleId", a.GetMenuTreeSelect)
}

// @Tags     菜单
// @Summary  Menu列表数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysMenuGetPageReq                 true "data"
// @Success  200  {object} ginx.ResponseData{data=[]schema.SysMenu} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                      "失败结果"
// @Router   /menu [get]
func (a *SysMenuApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysMenuGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysMenuSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     菜单
// @Summary  Menu详情数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id  path     uint64                                 false "id"
// @Success  200 {object} ginx.ResponseData{data=schema.SysMenu} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /menu/{id} [get]
func (a *SysMenuApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "id")

	result, err := a.SysMenuSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     菜单
// @Summary  创建菜单
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysMenuInsertReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                      "失败结果"
// @Router   /menu [post]
func (a *SysMenuApi) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysMenuInsertReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.CreateBy = contextx.FromUserID(ctx)

	err := a.SysMenuSrv.Insert(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     菜单
// @Summary  修改菜单
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysMenuUpdateReq true "body"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                                  "失败结果"
// @Router   /menu [put]
func (a *SysMenuApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysMenuUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysMenuSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     菜单
// @Summary  删除菜单
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysMenuDeleteReq true "body"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /menu [delete]
func (a *SysMenuApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysMenuDeleteReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	p := contextx.FromDataPermission(ctx)
	req.RoleID = p.RoleID

	err := a.SysMenuSrv.Delete(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     菜单
// @Summary  根据登录角色名称获取菜单列表数据（左菜单使用）
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Success  200 {object} ginx.ResponseData{data=[]schema.SysMenu} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /menurole [get]
func (a *SysMenuApi) GetMenuRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleName := contextx.FromRoleKey(ctx)
	result, err := a.SysMenuSrv.GetMenuRole(ctx, roleName)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     菜单
// @Summary  根据角色ID查询菜单下拉树结构
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    roleId path     uint64                                               true "roleId"
// @Success  200    {object} ginx.ResponseData{data=schema.GetMenuTreeSelectResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /roleMenuTreeselect/{roleId} [get]
func (a *SysMenuApi) GetMenuTreeSelect(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.SysMenuSrv.GetLabel(ctx)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	id := ginx.ParseParamID(c, "roleId")
	menuIds, err := a.SysRoleSrv.GetRoleMenuId(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, schema.GetMenuTreeSelectResp{
		Menus:       result,
		CheckedKeys: menuIds,
	})
}
