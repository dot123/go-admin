package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AuthApiSet = wire.NewSet(wire.Struct(new(AuthApi), "*"))

type AuthApi struct {
	AuthSrv *service.AuthSrv
}

// RegisterRoute 注册路由
func (a *AuthApi) RegisterRoute(r *gin.RouterGroup) {
	r.POST("login", a.Login)
	r.POST("refresh_token", a.RefreshToken)
	r.POST("logout", a.Logout)
}

// @Tags     认证
// @Summary 登陆
// @Accept   application/json
// @Produce  application/json
// @Param   data body     schema.LoginParam                             true "请求参数"
// @Success  200 {object} ginx.ResponseData{data=schema.LoginTokenInfo} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{} "失败结果"
// @Router  /login [post]
func (a *AuthApi) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.LoginParam
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.AuthSrv.Login(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     认证
// @Summary  刷新token
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Success 200  {object} ginx.ResponseData{data=schema.LoginTokenInfo} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                           "失败结果"
// @Router   /refresh_token [post]
func (a *AuthApi) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AuthSrv.DestroyToken(ctx, ginx.GetToken(c))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.AuthSrv.RefreshToken(ctx)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags    认证
// @Summary  登出
// @Accept  application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success  200 {object} ginx.ResponseData{} "成功结果"
// @Failure 500  {object} ginx.ResponseFail{}                           "失败结果"
// @Router   /logout [post]
func (a *AuthApi) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	if err := a.AuthSrv.DestroyToken(ctx, ginx.GetToken(c)); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
