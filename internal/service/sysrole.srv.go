package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var SysRoleSet = wire.NewSet(wire.Struct(new(SysRoleSrv), "*"))

type SysRoleSrv struct {
	SysRoleRepo *system.SysRoleRepo
	SysMenuRepo *system.SysMenuRepo
	SysDeptRepo *system.SysDeptRepo
	TransRepo   *util.Trans
	Enforcer    *casbin.SyncedEnforcer
}

func (s *SysRoleSrv) GetPage(ctx context.Context, req *schema.SysRoleGetPageReq) (*schema.SysRoleGetPageResp, error) {
	list, total, err := s.SysRoleRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	result := new(schema.SysRoleGetPageResp)
	result.Total = total
	copier.Copy(&result.Data, list)

	return result, err
}

func (s *SysRoleSrv) Get(ctx context.Context, id uint64) (*schema.SysRole, error) {
	model, err := s.SysRoleRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	result := new(schema.SysRole)
	copier.Copy(result, model)

	result.MenuIds, err = s.GetRoleMenuId(ctx, model.RoleID)
	if err != nil {
		logger.WithContext(ctx).Errorf("get menuIds error, %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (s *SysRoleSrv) GetRoleMenuId(ctx context.Context, roleId uint64) (*[]uint64, error) {
	model, err := s.SysRoleRepo.FindOneWithSysMenu(ctx, roleId)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}
	l := *model.SysMenu
	menuIds := make([]uint64, 0)
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuID)
	}
	return &menuIds, nil
}

func (s *SysRoleSrv) Insert(ctx context.Context, req *schema.SysRoleInsertReq) error {
	dataMenu, err := s.SysMenuRepo.FindAllWithApisByIds(ctx, &req.MenuIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	model := new(system.SysRole)

	copier.Copy(model, req)
	model.SysMenu = dataMenu

	role, err := s.SysRoleRepo.FindOneByRoleKey(ctx, req.RoleKey)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	if role.RoleKey == req.RoleKey {
		err = errors.NewDefaultResponse("roleKey已存在，需更换在提交！")
		return err
	}

	err = s.SysRoleRepo.Insert(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	mp := make(map[string]interface{}, 0)
	polices := make([][]string, 0)
	for _, menu := range *dataMenu {
		for _, api := range menu.SysApi {
			if mp[model.RoleKey+"-"+api.Path+"-"+api.Action] != "" {
				mp[model.RoleKey+"-"+api.Path+"-"+api.Action] = ""
				polices = append(polices, []string{model.RoleKey, api.Path, api.Action})
			}
		}
	}

	if len(polices) <= 0 {
		return nil
	}

	_, err = s.Enforcer.AddNamedPolicies("p", polices)
	return err
}

func (s *SysRoleSrv) Update(ctx context.Context, req *schema.SysRoleUpdateReq) error {
	model, err := s.SysRoleRepo.FindOneWithSysMenu(ctx, req.RoleID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.NewDefaultResponse("role角色不存在")
		}
		return err
	}

	mlist, err := s.SysMenuRepo.FindAllWithApisByIds(ctx, &req.MenuIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err := s.SysRoleRepo.DeleteSysMenu(ctx, model.RoleID, model.SysMenu)
		if err != nil {
			return err
		}

		copier.Copy(model, req)
		model.SysMenu = mlist

		err = s.SysRoleRepo.UpdateWithAssociation(ctx, model)
		if err != nil {
			return err
		}

		_, err = s.Enforcer.RemoveFilteredPolicy(0, model.RoleKey)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	mp := make(map[string]interface{}, 0)
	polices := make([][]string, 0)
	for _, menu := range *mlist {
		for _, api := range menu.SysApi {
			if mp[model.RoleKey+"-"+api.Path+"-"+api.Action] != "" {
				mp[model.RoleKey+"-"+api.Path+"-"+api.Action] = ""
				polices = append(polices, []string{model.RoleKey, api.Path, api.Action})
			}
		}
	}

	if len(polices) <= 0 {
		return nil
	}
	_, err = s.Enforcer.AddNamedPolicies("p", polices)

	return err
}

func (s *SysRoleSrv) Delete(ctx context.Context, id uint64) error {
	model, err := s.SysRoleRepo.FindOneWithSysMenuAndSysDept(ctx, id)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		rowsAffected, err := s.SysRoleRepo.Delete(ctx, model)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return err
		}
		if rowsAffected == 0 {
			return errors.NewDefaultResponse("无权更新该数据")
		}

		_, _ = s.Enforcer.RemoveFilteredPolicy(0, model.RoleKey)

		return nil
	})

	return nil
}

func (s *SysRoleSrv) UpdateDataScope(ctx context.Context, req *schema.RoleDataScopeReq) error {
	model, err := s.SysRoleRepo.FindOneWithSysDept(ctx, req.RoleID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	dlist, err := s.SysDeptRepo.FindAllByIds(ctx, &req.DeptIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err = s.SysRoleRepo.DeleteSysDept(ctx, model.RoleID, model.SysDept)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return err
		}

		copier.Copy(model, req)

		model.SysDept = dlist

		err = s.SysRoleRepo.UpdateWithAssociation(ctx, model)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return err
		}
		return nil
	})
	return err
}

func (s *SysRoleSrv) UpdateStatus(ctx context.Context, req *schema.UpdateStatusReq) error {
	err := s.SysRoleRepo.UpdateStatus(ctx, req.RoleID, req.Status, req.UpdateBy)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}
	return nil
}

func (s *SysRoleSrv) GetWithName(ctx context.Context, roleName string) (*schema.SysRole, error) {
	model, err := s.SysRoleRepo.FindOneByRoleName(ctx, roleName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.NewDefaultResponse("查看对象不存在或无权查看")
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}
	resp := new(schema.SysRole)

	copier.Copy(resp, model)

	menuIds, err := s.GetRoleMenuId(ctx, model.RoleID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	resp.MenuIds = menuIds
	return resp, nil
}

func (s *SysRoleSrv) GetById(ctx context.Context, roleId uint64) (*[]string, error) {
	permissions := make([]string, 0)
	model, err := s.SysRoleRepo.FindOneWithSysMenu(ctx, roleId)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return &permissions, nil
}
