package service

import (
	"context"
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sort"
	"strings"
)

var SysMenuSet = wire.NewSet(wire.Struct(new(SysMenuSrv), "*"))

type SysMenuSrv struct {
	SysMenuRepo *system.SysMenuRepo
	SysApiRepo  *system.SysApiRepo
	TransRepo   *util.Trans
	SysRoleRepo *system.SysRoleRepo
}

func (s *SysMenuSrv) GetPage(ctx context.Context, req *schema.SysMenuGetPageReq) (*[]*schema.SysMenu, error) {
	list, err := s.SysMenuRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	menu := make([]*schema.SysMenu, 0)
	for i := range *list {
		m := new(schema.SysMenu)
		copier.Copy(m, (*list)[i])
		menu = append(menu, m)
	}

	menus := make([]*schema.SysMenu, 0)
	for i := 0; i < len(menu); i++ {
		if menu[i].ParentID != 0 {
			continue
		}
		menusInfo := s.menuCall(&menu, menu[i])
		menus = append(menus, menusInfo)
	}
	return &menus, nil
}

func (s *SysMenuSrv) menuCall(menuList *[]*schema.SysMenu, menu *schema.SysMenu) *schema.SysMenu {
	list := *menuList

	min := make([]*schema.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuID != list[j].ParentID {
			continue
		}
		mi := new(schema.SysMenu)
		mi.MenuID = list[j].MenuID
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentID = list[j].ParentID
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.SysApi = list[j].SysApi
		mi.Children = make([]*schema.SysMenu, 0)

		if mi.MenuType != schema.Button {
			ms := s.menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func (s *SysMenuSrv) Get(ctx context.Context, id uint64) (*schema.SysMenu, error) {
	model, err := s.SysMenuRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.NewDefaultResponse("查看对象不存在或无权查看")
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	apis := make([]uint64, 0)
	for _, v := range model.SysApi {
		apis = append(apis, v.ID)
	}

	resp := new(schema.SysMenu)
	copier.Copy(resp, model)
	resp.Apis = apis

	return resp, nil
}

func (s *SysMenuSrv) Insert(ctx context.Context, req *schema.SysMenuInsertReq) error {
	model := new(system.SysMenu)
	copier.Copy(model, req)

	if err := s.SysMenuRepo.Insert(ctx, model); err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	err := s.initPaths(ctx, model)
	return err
}

func (s *SysMenuSrv) initPaths(ctx context.Context, menu *system.SysMenu) error {
	if menu.ParentID != 0 {
		parentMenu, err := s.SysMenuRepo.Get(ctx, menu.ParentID)
		if parentMenu.Paths == "" {
			err = errors.NewDefaultResponse("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return err
		}
		menu.Paths = fmt.Sprintf("%s/%d", parentMenu.Paths, menu.MenuID)
	} else {
		menu.Paths = fmt.Sprintf("/0/%d", menu.MenuID)
	}
	err := s.SysMenuRepo.UpdatePaths(ctx, menu.MenuID, menu.Paths)
	return err
}

func (s *SysMenuSrv) Update(ctx context.Context, req *schema.SysMenuUpdateReq) error {
	model, err := s.SysMenuRepo.Get(ctx, req.MenuID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	oldPath := model.Paths

	list, err := s.SysApiRepo.Find(ctx, &req.Apis)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	copier.Copy(model, req)

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := s.SysMenuRepo.DeleteSysApi(ctx, model); err != nil {
			return err
		}

		model.SysApi = *list

		if err := s.SysMenuRepo.UpdateWithAssociation(ctx, model); err != nil {
			return err
		}

		menuList, err := s.SysMenuRepo.FindByPath(ctx, oldPath)
		if err != nil {
			return err
		}

		for _, v := range *menuList {
			v.Paths = strings.Replace(v.Paths, oldPath, model.Paths, 1)
			s.SysMenuRepo.UpdatePaths(ctx, v.MenuID, v.Paths)
		}

		return nil
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	return nil
}

func (s *SysMenuSrv) Delete(ctx context.Context, req *schema.SysMenuDeleteReq) error {
	list, err := s.SysMenuRepo.FindAllByIds(ctx, &req.Ids)

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err := s.SysRoleRepo.DeleteSysMenu(ctx, req.RoleID, list)
		if err != nil {
			return err
		}

		rowsAffected, err := s.SysMenuRepo.Delete(ctx, list)
		if rowsAffected == 0 {
			err = errors.NewDefaultResponse("无权删除该数据")
			return err
		}

		return nil
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	return nil
}

func (s *SysMenuSrv) GetLabel(ctx context.Context) (*[]*schema.MenuLabel, error) {
	list, err := s.SysMenuRepo.FindAll(ctx)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	m := make([]*schema.MenuLabel, 0)

	for _, menu := range *list {
		if menu.ParentID != 0 {
			continue
		}
		e := new(schema.MenuLabel)
		e.ID = menu.MenuID
		e.Label = menu.Title
		deptsInfo := s.menuLabelCall(ctx, list, e)

		m = append(m, deptsInfo)
	}

	return &m, nil
}

func (s *SysMenuSrv) menuLabelCall(ctx context.Context, eList *[]*system.SysMenu, dept *schema.MenuLabel) *schema.MenuLabel {
	list := *eList

	min := make([]*schema.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.ID != list[j].ParentID {
			continue
		}
		mi := new(schema.MenuLabel)
		mi.ID = list[j].MenuID
		mi.Label = list[j].Title
		mi.Children = make([]*schema.MenuLabel, 0)
		if list[j].MenuType != schema.Button {
			ms := s.menuLabelCall(ctx, eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

func (s *SysMenuSrv) GetMenuRole(ctx context.Context, roleName string) (*[]*schema.SysMenu, error) {
	menus, err := s.getByRoleName(ctx, roleName)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	menu := make([]*schema.SysMenu, 0)
	for i := range *menus {
		m := new(schema.SysMenu)
		copier.Copy(m, (*menus)[i])
		menu = append(menu, m)
	}

	m := make([]*schema.SysMenu, 0)
	for i := 0; i < len(menu); i++ {
		if (*menus)[i].ParentID != 0 {
			continue
		}
		menusInfo := s.menuCall(&menu, menu[i])
		m = append(m, menusInfo)
	}
	return &m, nil
}

func (s *SysMenuSrv) getByRoleName(ctx context.Context, roleName string) (*[]*system.SysMenu, error) {
	if roleName == "admin" {
		list, err := s.SysMenuRepo.FindAllByMC(ctx)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return nil, err
		}
		sort.Sort(system.SysMenuSlice(*list))
		return list, nil
	} else {
		role, err := s.SysRoleRepo.FindOneWithSysMenuByRoleKey(ctx, roleName)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return nil, err
		}
		if role.SysMenu != nil {
			ids := make([]uint64, 0)
			for _, menu := range *role.SysMenu {
				ids = append(ids, menu.MenuID)
			}
			menus := make([]*system.SysMenu, 0)
			if err := s.recursiveSetMenu(ctx, &ids, &menus); err != nil {
				return nil, err
			}

			menu := s.menuDistinct(&menus)

			sort.Sort(system.SysMenuSlice(*menu))
			return menu, nil
		}
	}

	return &[]*system.SysMenu{}, nil
}

func (s *SysMenuSrv) menuDistinct(menuList *[]*system.SysMenu) *[]*system.SysMenu {
	result := make([]*system.SysMenu, 0)
	distinctMap := make(map[uint64]struct{}, len(*menuList))
	for _, menu := range *menuList {
		if _, ok := distinctMap[menu.MenuID]; !ok {
			distinctMap[menu.MenuID] = struct{}{}
			result = append(result, menu)
		}
	}
	return &result
}

func (s *SysMenuSrv) recursiveSetMenu(ctx context.Context, ids *[]uint64, menus *[]*system.SysMenu) error {
	if len(*ids) == 0 || menus == nil {
		return nil
	}

	subMenus, err := s.SysMenuRepo.FindAllByIds(ctx, ids)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	subIds := make([]uint64, 0)
	for _, menu := range *subMenus {
		if menu.ParentID != 0 {
			subIds = append(subIds, menu.ParentID)
		}
		if menu.MenuType != schema.Button {
			*menus = append(*menus, menu)
		}
	}
	return s.recursiveSetMenu(ctx, &subIds, menus)
}
