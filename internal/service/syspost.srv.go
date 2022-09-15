package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var SysPostSet = wire.NewSet(wire.Struct(new(SysPostSrv), "*"))

type SysPostSrv struct {
	SysPostRepo *system.SysPostRepo
}

func (s *SysPostSrv) GetPage(ctx context.Context, req *schema.SysPostPageReq) (*schema.SysPostPageResp, error) {
	list, total, err := s.SysPostRepo.GetPage(ctx, req)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	result := new(schema.SysPostPageResp)
	result.Total = total
	copier.Copy(&result.Data, list)

	return result, err
}

func (s *SysPostSrv) Get(ctx context.Context, id uint64) (*schema.SysPost, error) {
	model, err := s.SysPostRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.NewDefaultResponse("查看对象不存在或无权查看")
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, err
	}

	result := new(schema.SysPost)
	copier.Copy(result, model)

	return result, nil
}

func (s *SysPostSrv) Insert(ctx context.Context, req *schema.SysPostInsertReq) error {
	model := new(system.SysPost)
	copier.Copy(model, req)

	err := s.SysPostRepo.Insert(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}
	return nil
}

func (s *SysPostSrv) Update(ctx context.Context, req *schema.SysPostUpdateReq) error {
	model, err := s.SysPostRepo.Get(ctx, req.PostID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("无权更新该数据")
	}

	copier.Copy(model, req)

	err = s.SysPostRepo.Update(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}
	return err
}

func (s *SysPostSrv) Delete(ctx context.Context, req *schema.SysPostDeleteReq) error {
	rowsAffected, err := s.SysPostRepo.Delete(ctx, &req.Ids)
	if rowsAffected == 0 {
		return errors.NewDefaultResponse("无权更新该数据")
	}
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return err
	}
	return nil
}
