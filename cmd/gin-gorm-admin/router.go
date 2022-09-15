package main

import (
	"github.com/casbin/casbin/v2"
	v1 "github.com/dot123/gin-gorm-admin/api/v1"
	"github.com/dot123/gin-gorm-admin/internal/middleware"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	Enforcer   *casbin.SyncedEnforcer
	AuthSrv    *service.AuthSrv
	AuthApi    *v1.AuthApi
	MonitorApi *v1.MonitorApi
	FileApi    *v1.FileApi
	MsgApi     *v1.MsgApi
	SysApi     *v1.SysApi
	SysUserApi *v1.SysUserApi
	SysRoleApi *v1.SysRoleApi
	SysPostApi *v1.SysPostApi
	SysDeptApi *v1.SysDeptApi
	SysMenuApi *v1.SysMenuApi
}

func (m *Router) Register(app *gin.Engine) error {
	m.RegisterAPI(app)
	return nil
}

func (m *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI 注册路由
func (m *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(middleware.UserAuthMiddleware(m.AuthSrv, middleware.AllowPathPrefixSkipper("/api/v1/login")))
	g.Use(middleware.CasbinMiddleware(m.Enforcer, middleware.AllowPathPrefixSkipper("/api/v1/login")))

	v1 := g.Group("/v1")
	m.AuthApi.RegisterRoute(v1)
	m.MonitorApi.RegisterRoute(v1)
	m.FileApi.RegisterRoute(v1)
	m.MsgApi.RegisterRoute(v1)

	m.SysApi.RegisterRoute(v1)
	m.SysUserApi.RegisterRoute(v1)
	m.SysRoleApi.RegisterRoute(v1)
	m.SysPostApi.RegisterRoute(v1)
	m.SysDeptApi.RegisterRoute(v1)
	m.SysMenuApi.RegisterRoute(v1)
}
