package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(e *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		ctx := c.Request.Context()
		roleKey := contextx.FromRoleKey(ctx)

		casbinExclude := false
		for _, i := range CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			logger.WithContext(c).Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		res, err := e.Enforce(roleKey, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			logger.WithContext(c).Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			ginx.ResError(c, err)
			return
		}

		if res {
			logger.WithContext(c).Infof("isTrue: %v role: %s method: %s path: %s", res, roleKey, c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			logger.WithContext(c).Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, roleKey, c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
