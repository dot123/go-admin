package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"net/http"
)

var SysUserApiSet = wire.NewSet(wire.Struct(new(SysUserApi), "*"))

type SysUserApi struct {
	SysUserSrv *service.SysUserSrv
	SysRoleSrv *service.SysRoleSrv
}

// RegisterRoute 注册路由
func (a *SysUserApi) RegisterRoute(r *gin.RouterGroup) {
	r.GET("sys-user", a.GetPage)
	r.GET("sys-user/:userId", a.Get)
	r.POST("sys-user", a.Insert)
	r.PUT("sys-user", a.Update)
	r.DELETE("sys-user", a.Delete)
	r.POST("user/avatar", a.InsetAvatar)
	r.PUT("user/status", a.UpdateStatus)
	r.PUT("user/pwd/reset", a.ResetPwd)
	r.PUT("user/pwd/set", a.UpdatePwd)
	r.GET("user/profile/:userId", a.GetProfile)
	r.GET("user/getinfo", a.GetInfo)
}

// @Tags     用户
// @Summary  列表用户信息数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.SysUserGetPageReq                          true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=schema.SysUserGetPageResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                               "失败结果"
// @Router   /sys-user [get]
func (a *SysUserApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysUserGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysUserSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     用户
// @Summary  获取用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    userId path     uint64                                 true "userId"
// @Success  200    {object} ginx.ResponseData{data=schema.SysUser} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /sys-user/{userId} [get]
func (a *SysUserApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "userId")

	result, err := a.SysUserSrv.Get(ctx, id)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     用户
// @Summary  创建用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysUserInsertReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{} "失败结果"
// @Router   /sys-user [post]
func (a *SysUserApi) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysUserInsertReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.CreateBy = contextx.FromUserID(ctx)

	err := a.SysUserSrv.Insert(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     用户
// @Summary  修改用户数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysUserUpdateReq true "请求参数"
// @Success  200 {object} ginx.ResponseData{} "成功结果"
// @Failure  500    {object} ginx.ResponseFail{}                           "失败结果"
// @Router   /sys-user [put]
func (a *SysUserApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysUserUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysUserSrv.Update(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     用户
// @Summary  删除用户数据
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.SysUserDeleteReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500 {object} ginx.ResponseFail{} "失败结果"
// @Router   /sys-user [delete]
func (a *SysUserApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysUserDeleteReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SysUserSrv.Delete(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     个人中心
// @Summary  修改头像
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Accept   multipart/form-data
// @Param    file formData file                true "file"
// @Success  200  {object} ginx.ResponseData{}       "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}       "失败结果"
// @Router   /user/avatar [post]
func (a *SysUserApi) InsetAvatar(c *gin.Context) {
	ctx := c.Request.Context()

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		logger.WithContext(ctx).Errorf("upload avatar file: %s", file.Filename)
		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, filPath)
		if err != nil {
			logger.WithContext(ctx).Errorf("save file error, %s", err.Error())
			ginx.ResError(c, err, 500)
			return
		}
	}

	var req schema.UpdateSysUserAvatarReq
	req.UserID = contextx.FromUserID(ctx)
	req.Avatar = "/" + filPath
	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysUserSrv.UpdateAvatar(ctx, &req)
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return
	}
	ginx.ResData(c, filPath)
}

// @Tags     个人中心
// @Summary  修改用户状态
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.UpdateSysUserStatusReq true "body"
// @Success  200  {object} ginx.ResponseData{}           "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}           "失败结果"
// @Router   /user/status [put]
func (a *SysUserApi) UpdateStatus(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.UpdateSysUserStatusReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysUserSrv.UpdateStatus(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     用户
// @Summary  重置用户密码
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.ResetSysUserPwdReq true "body"
// @Success  200  {object} ginx.ResponseData{} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{} "失败结果"
// @Router   /user/pwd/reset [put]
func (a *SysUserApi) ResetPwd(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.ResetSysUserPwdReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	req.UpdateBy = contextx.FromUserID(ctx)

	err := a.SysUserSrv.ResetPwd(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     用户
// @Summary  修改密码
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.PassWord     true "body"
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /user/pwd/set [put]
func (a *SysUserApi) UpdatePwd(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.PassWord
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SysUserSrv.UpdatePwd(ctx, contextx.FromUserID(ctx), req.OldPassword, req.NewPassword)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     个人中心
// @Summary  获取个人中心用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    userId path     uint64                                        false "userId"
// @Success  200    {object} ginx.ResponseData{data=schema.GetProfileResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /user/profile/{userId} [get]
func (a *SysUserApi) GetProfile(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.SysUserById
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	id := ginx.ParseParamID(c, "userId")

	user, roles, posts, err := a.SysUserSrv.GetProfile(ctx, id)
	if err != nil {
		logger.WithContext(ctx).Errorf("get user profile error, %s", err.Error())
		ginx.ResError(c, err, 500)
		return
	}

	ginx.ResData(c, &schema.GetProfileResp{
		User:  user,
		Roles: roles,
		Posts: posts,
	})
}

// @Tags     个人中心
// @Summary  获取个人信息
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Success  200  {object} ginx.ResponseData{}     "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}     "失败结果"
// @Router   /user/getinfo [get]
func (a *SysUserApi) GetInfo(c *gin.Context) {
	ctx := c.Request.Context()

	var roles = make([]string, 1)
	roles[0] = contextx.FromRoleKey(ctx)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	var mp = make(map[string]interface{})
	mp["roles"] = roles

	if contextx.FromRoleKey(ctx) == "admin" || contextx.FromRoleKey(ctx) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := a.SysRoleSrv.GetById(ctx, contextx.FromUserID(ctx))
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysUser, err := a.SysUserSrv.Get(ctx, contextx.FromUserID(ctx))
	if err != nil {
		ginx.ResError(c, err, http.StatusUnauthorized)
		return
	}
	mp["introduction"] = " am a super administrator"
	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if sysUser.Avatar != "" {
		mp["avatar"] = sysUser.Avatar
	}
	mp["userName"] = sysUser.NickName
	mp["userId"] = sysUser.UserID
	mp["deptId"] = sysUser.DeptID
	mp["name"] = sysUser.NickName
	mp["code"] = 200
	ginx.ResData(c, mp)
}
