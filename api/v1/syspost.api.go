package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SysPostApiSet = wire.NewSet(wire.Struct(new(SysPostApi), "*"))

type SysPostApi struct {
	SysPostSrv *service.SysPostSrv
}

// RegisterRoute 注册路由
func (a *SysPostApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("post", a.GetPage)
	r.GET("post/:postId", a.Get)
	r.POST("post", a.Insert)
	r.PUT("post", a.Update)
	r.DELETE("post", a.Delete)
}

// @Tags     岗位
// @Summary  岗位列表数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysPostPageReq                          true "body"
// @Success  200  {object} ginx.ResponseData{data=schema.SysPostPageResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                            "失败结果"
// @Router   /post [get]
func (a *SysPostApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysPostPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysPostSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     岗位
// @Summary  获取岗位信息
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    postId path     uint64                                 true "编码"
// @Success  200    {object} ginx.ResponseData{data=schema.SysPost} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /post/{postId} [get]
func (a *SysPostApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "postId")
	result, err := a.SysPostSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     岗位
// @Summary  添加岗位
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysPostInsertReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /post [post]
func (a *SysPostApi) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysPostInsertReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SysPostSrv.Insert(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     岗位
// @Summary  修改岗位
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysPostUpdateReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /post [put]
func (a *SysPostApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysPostUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysPostSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     岗位
// @Summary  删除岗位
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysPostDeleteReq true "data"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /post [delete]
func (a *SysPostApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysPostDeleteReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SysPostSrv.Delete(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
