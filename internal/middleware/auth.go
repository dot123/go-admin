package middleware

import (
	"encoding/json"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/dot123/gin-gorm-admin/pkg/helper"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
)

func wrapUserAuthContext(c *gin.Context, userID uint64, username, roleKey string, p *schema.DataPermission) {
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = contextx.NewUserName(ctx, username)
	ctx = contextx.NewRoleKey(ctx, roleKey)
	ctx = contextx.NewDataPermission(ctx, p)
	ctx = logger.NewUserIDContext(ctx, userID)
	ctx = logger.NewUserNameContext(ctx, username)
	c.Request = c.Request.WithContext(ctx)
}

func UserAuthMiddleware(s *service.AuthSrv, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		data, err := s.ParseUserData(ginx.GetToken(c))
		if err != nil {
			ginx.ResError(c, err)
			return
		}

		claims := *data
		userID := uint64(claims["userid"].(float64))
		username := claims["username"].(string)
		roleKey := claims["rolekey"].(string)
		jsonDataPermission := claims["dataPermission"].(string)

		p := new(schema.DataPermission)
		json.Unmarshal(helper.S2B(jsonDataPermission), p)

		wrapUserAuthContext(c, userID, username, roleKey, p)

		c.Next()
	}
}
