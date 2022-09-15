package schema

import (
	"github.com/dot123/gin-gorm-admin/internal/validate"
)

type LoginParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
	UUID     string `json:"uuid" validate:"required"`
}

func (m *LoginParam) Verify() string {
	messages := map[string]string{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
		"Code.required":     "验证码不能为空",
		"UUID.required":     "uuid不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type LoginTokenInfo struct {
	AccessToken  string `json:"access_token"`  // 访问令牌
	AccessExpire int64  `json:"access_expire"` // 过期时间戳
	RefreshAfter int64  `json:"refresh_after"` // 刷新时间戳
}
