package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/helper"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/dot123/gin-gorm-admin/pkg/redisHelper"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func NewAuthSrv(sysUserRepo *system.SysUserRepo, sysRoleRepo *system.SysRoleRepo) *AuthSrv {
	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       0,
	})
	return &AuthSrv{SysUserRepo: sysUserRepo, SysRoleRepo: sysRoleRepo, Ring: ring}
}

type AuthSrv struct {
	SysUserRepo *system.SysUserRepo
	SysRoleRepo *system.SysRoleRepo
	Ring        *redis.Ring
}

func (s *AuthSrv) Login(ctx context.Context, params *schema.LoginParam) (*schema.LoginTokenInfo, error) {
	user, err := s.SysUserRepo.GetUserByName(ctx, params.Username)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		logger.WithContext(ctx).Errorf("用户不存在,参数:%s,异常:%s", params.Username, err.Error())
		return nil, errors.NewDefaultResponse("用户不存在")
	default:
		logger.WithContext(ctx).Errorf("用户登录失败,参数:%s,异常:%s", params.Username, err.Error())
		return nil, errors.NewDefaultResponse("用户登录失败")
	}

	if user.Status != 2 {
		return nil, errors.NewDefaultResponse("用户已停用")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		logger.WithContext(ctx).Errorf("用户密码不正确,参数:%s", params.Password)
		return nil, errors.NewDefaultResponse("用户密码不正确")
	}
	role, err := s.SysRoleRepo.Get(ctx, user.RoleID)
	if err != nil {
		return nil, errors.NewDefaultResponse("获取角色失败")
	}

	return s.GenerateToken(ctx, user.UserID, user.DeptID, role.RoleID, role.DataScope, user.Username, role.RoleKey)
}

func (s *AuthSrv) GenerateToken(ctx context.Context,
	userID, deptID, roleID uint64, dataScope, username, roleKey string) (*schema.LoginTokenInfo, error) {
	now := time.Now().Unix()
	accessExpire := config.C.JWTAuth.AccessExpire

	p := schema.DataPermission{
		UserID:    userID,
		DeptID:    deptID,
		RoleID:    roleID,
		DataScope: dataScope,
	}

	jsonDataPermission, _ := json.Marshal(p)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":            now + accessExpire,
		"iat":            now,
		"userid":         userID,
		"username":       username,
		"rolekey":        roleKey,
		"dataPermission": helper.B2S(jsonDataPermission),
	})
	jwtToken, err := token.SignedString([]byte(config.C.JWTAuth.AccessSecret))

	if err != nil {
		logger.WithContext(ctx).Errorf("生成token失败,参数:%d,%d,%d,%s,%s,%s,异常:%s",
			userID, deptID, roleID, dataScope, username, roleKey, err.Error())
		return nil, err
	}

	item := &schema.LoginTokenInfo{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	return item, nil
}

func (s *AuthSrv) RefreshToken(ctx context.Context) (*schema.LoginTokenInfo, error) {
	p := contextx.FromDataPermission(ctx)
	return s.GenerateToken(ctx, p.UserID, p.DeptID, p.RoleID, p.DataScope, contextx.FromUserName(ctx), contextx.FromRoleKey(ctx))
}

// DestroyToken 销毁令牌
func (s *AuthSrv) DestroyToken(ctx context.Context, tokenString string) error {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		logger.WithContext(ctx).Errorf("destroyToken error:%s", err.Error())
		return errors.NewDefaultResponse("销毁令牌失败")
	}

	exp := int64((*claims)["exp"].(float64))
	expired := time.Unix(exp, 0).Sub(time.Now())
	if err = redisHelper.Set(s.Ring, tokenString, "1", expired); err != nil {
		logger.WithContext(ctx).Errorf("destroyToken error:%s", err.Error())
		return errors.NewDefaultResponse("销毁令牌失败")
	}
	return nil
}

func (s *AuthSrv) parseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.C.JWTAuth.AccessSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return &claims, nil
}

func (s *AuthSrv) ParseUserData(tokenString string) (*jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, errors.NewDefaultResponse("无效的token")
	}

	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if exists, err := redisHelper.Check(s.Ring, tokenString); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.NewDefaultResponse("无效的token")
	}
	return claims, nil
}
