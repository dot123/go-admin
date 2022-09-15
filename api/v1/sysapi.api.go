package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SysApiSet = wire.NewSet(wire.Struct(new(SysApi), "*"))

type SysApi struct {
	SysApiSrv *service.SysApiSrv
}

// RegisterRoute 注册路由
func (a *SysApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("sys-api", a.GetPage)
	r.GET("sys-api/:id", a.Get)
	r.PUT("sys-api", a.Update)
	r.DELETE("sys-api", a.Delete)
}

// @Tags     接口管理
// @Summary  获取接口管理列表
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysApiGetPageReq                          true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=schema.SysApiGetPageResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                              "失败结果"
// @Router   /sys-api [get]
func (a *SysApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysApiGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysApiSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     接口管理
// @Summary  获取接口管理
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id  path     uint64                                true "id"
// @Success  200 {object} ginx.ResponseData{data=schema.SysApi} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                   "失败结果"
// @Router   /sys-api/{id} [get]
func (a *SysApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "id")

	result, err := a.SysApiSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     接口管理
// @Summary  修改接口管理
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysApiUpdateReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}    "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}    "失败结果"
// @Router   /sys-api [put]
func (a *SysApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysApiUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	// 设置修改人id
	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysApiSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     接口管理
// @Summary  删除接口管理
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysApiDeleteReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}    "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}    "失败结果"
// @Router   /sys-api [delete]
func (a *SysApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysApiDeleteReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SysApiSrv.Delete(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
