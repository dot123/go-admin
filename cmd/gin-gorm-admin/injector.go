package main

import (
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine    *gin.Engine
	SysApiSrv *service.SysApiSrv
}
