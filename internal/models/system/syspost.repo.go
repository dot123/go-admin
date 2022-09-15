package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var SysPostSet = wire.NewSet(wire.Struct(new(SysPostRepo), "*"))

type SysPostRepo struct {
	DB *gorm.DB
}

func (m *SysPostRepo) GetPage(ctx context.Context, req *schema.SysPostPageReq) (*[]*SysPost, int64, error) {
	db := GetSysPostDB(ctx, m.DB)

	if req.PostID > 0 {
		db = db.Where("post_id = ?", req.PostID)
	}

	if req.PostName != "" {
		db = db.Where("post_name = ?", req.PostName)
	}

	if req.PostCode != "" {
		db = db.Where("post_code = ?", req.PostCode)
	}

	if req.Sort > 0 {
		db = db.Where("sort = ?", req.Sort)
	}

	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}

	if req.Remark != "" {
		db = db.Where("remark = ?", req.Remark)
	}

	list := make([]*SysPost, 0)
	total, err := util.GetPages(db, &list, req.PageNum, req.PageSize)
	return &list, total, err
}

func (m *SysPostRepo) Get(ctx context.Context, id uint64) (*SysPost, error) {
	model := new(SysPost)
	err := GetSysPostDB(ctx, m.DB).First(model, id).Error
	return model, err
}

func (m *SysPostRepo) Insert(ctx context.Context, model *SysPost) error {
	err := GetSysPostDB(ctx, m.DB).Create(model).Error
	return err
}

func (m *SysPostRepo) Update(ctx context.Context, model *SysPost) error {
	err := m.DB.Save(model).Error
	return err
}

func (m *SysPostRepo) Delete(ctx context.Context, ids *[]uint64) (rowsAffected int64, err error) {
	db := GetSysPostDB(ctx, m.DB).Delete(new(SysPost), ids)
	return db.RowsAffected, db.Error
}

func (m *SysPostRepo) FindAllByIds(ctx context.Context, ids *[]uint64) (*[]*SysPost, error) {
	list := make([]*SysPost, 0)
	err := GetSysPostDB(ctx, m.DB).Find(&list, ids).Error
	return &list, err
}
