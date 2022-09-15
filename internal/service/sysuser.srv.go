package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var SysUserSet = wire.NewSet(wire.Struct(new(SysUserSrv), "*"))

type SysUserSrv struct {
	SysUserRepo *system.SysUserRepo
	SysRoleRepo *system.SysRoleRepo
	SysPostRepo *system.SysPostRepo
}

func (s *SysUserSrv) GetPage(ctx context.Context, req *schema.SysUserGetPageReq) (*schema.SysUserGetPageResp, error) {
	list, total, err := s.SysUserRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取用户列表失败")
	}

	result := new(schema.SysUserGetPageResp)
	result.Total = total
	copier.Copy(&result.Data, list)

	return result, err
}

func (s *SysUserSrv) Get(ctx context.Context, id uint64) (*schema.SysUser, error) {
	model, err := s.SysUserRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("查看对象不存在或无权查看")
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取用户失败")
	}

	result := new(schema.SysUser)
	copier.Copy(result, model)

	return result, nil
}

func (s *SysUserSrv) Insert(ctx context.Context, req *schema.SysUserInsertReq) error {
	user, err := s.SysUserRepo.GetUserByName(ctx, req.Username)
	if user.Username == req.Username {
		return errors.NewDefaultResponse("用户名已存在！")
	}

	model := new(system.SysUser)
	copier.Copy(model, req)

	err = s.SysUserRepo.Insert(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("新建用户失败")
	}
	return nil
}

func (s *SysUserSrv) Update(ctx context.Context, req *schema.SysUserUpdateReq) error {
	model, err := s.SysUserRepo.Get(ctx, req.UserID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.New("无权更新该数据")
	}

	copier.Copy(model, req)

	err = s.SysUserRepo.Update(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新用户信息失败")
	}
	return err
}

func (s *SysUserSrv) UpdateAvatar(ctx context.Context, req *schema.UpdateSysUserAvatarReq) error {
	if _, err := s.SysUserRepo.Get(ctx, req.UserID); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("无权更新该数据")
	}

	if err := s.SysUserRepo.UpdateAvatar(ctx, req.UserID, req.Avatar, req.UpdateBy); err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新头像失败")
	}
	return nil
}

func (s *SysUserSrv) UpdateStatus(ctx context.Context, req *schema.UpdateSysUserStatusReq) error {
	_, err := s.SysUserRepo.Get(ctx, req.UserID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("无权更新该数据")
	}

	err = s.SysUserRepo.UpdateStatus(ctx, req.UserID, req.Status, req.UpdateBy)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新用户状态失败")
	}
	return nil
}

func (s *SysUserSrv) ResetPwd(ctx context.Context, req *schema.ResetSysUserPwdReq) error {
	_, err := s.SysUserRepo.Get(ctx, req.UserID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("无权更新该数据")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.WithContext(ctx).Errorf("密码加密失败:%s", err.Error())
		return err
	}

	err = s.SysUserRepo.ResetPwd(ctx, req.UserID, string(hash), req.UpdateBy)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("重置用户密码失败")
	}
	return nil
}

func (s *SysUserSrv) Delete(ctx context.Context, req *schema.SysUserDeleteReq) error {
	rowsAffected, err := s.SysUserRepo.Delete(ctx, &req.Ids)
	if rowsAffected == 0 {
		return errors.NewDefaultResponse("无权更新该数据")
	}
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除用户失败")
	}
	return nil
}

func (s *SysUserSrv) UpdatePwd(ctx context.Context, id uint64, oldPassword, newPassword string) error {
	if newPassword == "" {
		return nil
	}

	model, err := s.SysUserRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.NewDefaultResponse("无权更新该数据")
		}
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(oldPassword)); err != nil {
		return errors.NewDefaultResponse("密码不正确")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.WithContext(ctx).Errorf("密码加密失败:%s", err.Error())
		return err
	}

	rowsAffected, err := s.SysUserRepo.UpdatePwd(ctx, id, string(hash))
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	if rowsAffected == 0 {
		err = errors.New("set password error")
		return err
	}
	return nil
}

func (s *SysUserSrv) GetProfile(ctx context.Context, userID uint64) (*schema.SysUser, *[]*schema.SysRole, *[]*schema.SysPost, error) {
	sysUser, err := s.SysUserRepo.FindOneWithDeptById(ctx, userID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, nil, nil, err
	}

	sysRole, err := s.SysRoleRepo.Get(ctx, sysUser.RoleID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, nil, nil, err
	}

	sysPosts, err := s.SysPostRepo.FindAllByIds(ctx, &sysUser.PostIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, nil, nil, err
	}

	user := new(schema.SysUser)
	roles := make([]*schema.SysRole, 0)
	posts := make([]*schema.SysPost, 0)

	copier.Copy(user, sysUser)
	copier.Copy(&posts, sysPosts)
	role := new(schema.SysRole)
	copier.Copy(role, sysRole)
	roles = append(roles, role)

	return user, &roles, &posts, nil
}
