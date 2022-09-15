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
)

var SysDeptSet = wire.NewSet(wire.Struct(new(SysDeptSrv), "*"))

type SysDeptSrv struct {
	SysDeptRepo *system.SysDeptRepo
	SysRoleRepo *system.SysRoleRepo
	SysUserRepo *system.SysUserRepo
	TransRepo   *util.Trans
}

func (s *SysDeptSrv) Get(ctx context.Context, id uint64) (*schema.SysDept, error) {
	model, err := s.SysDeptRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.NewDefaultResponse("查看对象不存在或无权查看")
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	result := new(schema.SysDept)
	copier.Copy(result, model)

	return result, nil
}

func (s *SysDeptSrv) Insert(ctx context.Context, req *schema.SysDeptInsertReq) error {
	model := new(system.SysDept)
	copier.Copy(model, req)

	err := s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := s.SysDeptRepo.Insert(ctx, model); err != nil {
			return err
		}

		deptPath := fmt.Sprintf("%d/", model.DeptID)
		if model.ParentID != 0 {
			deptP, err := s.SysDeptRepo.Get(ctx, model.ParentID)
			if err != nil {
				return err
			}
			deptPath = deptP.DeptPath + deptPath
		} else {
			deptPath = "/0/" + deptPath
		}

		if err := s.SysDeptRepo.UpdateDeptPath(ctx, model.DeptID, deptPath); err != nil {
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

func (s *SysDeptSrv) Update(ctx context.Context, req *schema.SysDeptUpdateReq) error {
	model, err := s.SysDeptRepo.Get(ctx, req.DeptID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	copier.Copy(model, req)

	deptPath := fmt.Sprintf("%d/", model.DeptID)
	if model.ParentID != 0 {
		deptP, err := s.SysDeptRepo.Get(ctx, model.ParentID)
		if err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return err
		}
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	model.DeptPath = deptPath

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		s.SysDeptRepo.Update(ctx, model)
		if err != nil {
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

func (s *SysDeptSrv) Delete(ctx context.Context, req *schema.SysDeptDeleteReq) error {
	list, err := s.SysDeptRepo.FindAllByIds(ctx, &req.Ids)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		s.SysUserRepo.DeleteSysDept(ctx, req.UserID, list)
		if err != nil {
			return err
		}

		s.SysRoleRepo.DeleteSysDept(ctx, req.RoleID, list)
		if err != nil {
			return err
		}

		rowsAffected, err := s.SysDeptRepo.Delete(ctx, list)
		if rowsAffected == 0 {
			err = errors.NewDefaultResponse("无权删除该数据")
			return err
		}

		if err != nil {
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

func (s *SysDeptSrv) GetPage(ctx context.Context, req *schema.SysDeptGetPageReq) (*[]*schema.DeptLabel, error) {
	list, err := s.SysDeptRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}
	m := make([]*schema.DeptLabel, 0)

	for _, dept := range *list {
		if dept.ParentID != 0 {
			continue
		}
		e := new(schema.DeptLabel)
		e.ID = dept.DeptID
		e.Label = dept.DeptName
		deptsInfo := s.deptTreeCall(ctx, list, e)

		m = append(m, deptsInfo)
	}

	return &m, nil
}

func (s *SysDeptSrv) deptTreeCall(ctx context.Context, deptList *[]*system.SysDept, dept *schema.DeptLabel) *schema.DeptLabel {
	list := *deptList
	min := make([]*schema.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.ID != list[j].ParentID {
			continue
		}
		mi := schema.DeptLabel{ID: list[j].DeptID, Label: list[j].DeptName, Children: []*schema.DeptLabel{}}
		ms := s.deptTreeCall(ctx, deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

func (s *SysDeptSrv) deptPageCall(ctx context.Context, deptlist *[]*system.SysDept, menu *schema.SysDept) *schema.SysDept {
	list := *deptlist
	min := make([]*schema.SysDept, 0)
	for j := 0; j < len(list); j++ {
		if menu.DeptID != list[j].ParentID {
			continue
		}
		mi := new(schema.SysDept)
		mi.DeptID = list[j].DeptID
		mi.ParentID = list[j].ParentID
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = make([]*schema.SysDept, 0)
		ms := s.deptPageCall(ctx, deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

func (s *SysDeptSrv) GetWithRoleId(ctx context.Context, roleId uint64) (*[]uint64, error) {
	deptIds := make([]uint64, 0)
	deptList, err := s.SysDeptRepo.GetWithRoleId(ctx, roleId)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	for _, roleDept := range *deptList {
		deptIds = append(deptIds, roleDept.DeptID)
	}
	return &deptIds, nil
}

func (s *SysDeptSrv) GetDeptLabel(ctx context.Context) (*[]*schema.DeptLabel, error) {
	list, err := s.SysDeptRepo.FindAll(ctx)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	m := make([]*schema.DeptLabel, 0)
	for _, dept := range *list {
		if dept.ParentID != 0 {
			continue
		}
		item := new(schema.DeptLabel)
		item.ID = dept.DeptID
		item.Label = dept.DeptName
		deptInfo := s.deptLabelCall(list, item)
		m = append(m, deptInfo)
	}
	return &m, nil
}

func (s *SysDeptSrv) deptLabelCall(deptList *[]*system.SysDept, dept *schema.DeptLabel) *schema.DeptLabel {
	list := *deptList
	min := make([]*schema.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.ID != list[j].ParentID {
			continue
		}
		mi := schema.DeptLabel{ID: list[j].DeptID, Label: list[j].DeptName, Children: make([]*schema.DeptLabel, 0)}
		ms := s.deptLabelCall(deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

func (s *SysDeptSrv) GetDeptTree(ctx context.Context, req *schema.SysDeptGetPageReq) (*[]*schema.DeptLabel, error) {
	list, err := s.SysDeptRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	m := make([]*schema.DeptLabel, 0)

	for _, dept := range *list {
		if dept.ParentID != 0 {
			continue
		}
		e := new(schema.DeptLabel)
		e.ID = dept.DeptID
		e.Label = dept.DeptName
		deptsInfo := s.deptTreeCall(ctx, list, e)

		m = append(m, deptsInfo)
	}
	return &m, nil
}
