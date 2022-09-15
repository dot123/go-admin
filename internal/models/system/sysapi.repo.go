package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var SysApiSet = wire.NewSet(wire.Struct(new(SysApiRepo), "*"))

type SysApiRepo struct {
	DB *gorm.DB
}

func (m *SysApiRepo) GetPage(ctx context.Context, req *schema.SysApiGetPageReq) (*[]*SysApi, int64, error) {
	db := GetSysApiDB(ctx, m.DB)
	if req.Title != "" {
		db = db.Where("title = ?", req.Title)
	}

	if req.Path != "" {
		db = db.Where("path = ?", req.Path)
	}

	if req.Action != "" {
		db = db.Where("action = ?", req.Action)
	}

	if req.TitleOrder != "" {
		db = db.Order("title " + req.TitleOrder)
	}

	if req.PathOrder != "" {
		db = db.Order("path " + req.PathOrder)
	}

	if req.CreatedAtOrder != "" {
		db = db.Order("created_at " + req.CreatedAtOrder)
	}

	db = db.Scopes(Permission(ctx, "sys_api"))

	list := make([]*SysApi, 0)
	total, err := util.GetPages(db, &list, req.PageNum, req.PageSize)
	return &list, total, err
}

func (m *SysApiRepo) Get(ctx context.Context, id uint64) (*SysApi, error) {
	data := new(SysApi)
	err := GetSysApiDB(ctx, m.DB).Scopes(Permission(ctx, "sys_api")).First(data, id).Error
	return data, err
}

func (m *SysApiRepo) Update(ctx context.Context, model *SysApi) error {
	err := GetSysApiDB(ctx, m.DB).Where("`id`=?", model.ID).Save(model).Error
	return err
}

func (m *SysApiRepo) Delete(ctx context.Context, ids *[]uint64) (int64, error) {
	db := GetSysApiDB(ctx, m.DB).Scopes(Permission(ctx, "sys_api")).Delete(new(SysApi), ids)
	return db.RowsAffected, db.Error
}

func (m *SysApiRepo) Find(ctx context.Context, ids *[]uint64) (*[]*SysApi, error) {
	list := make([]*SysApi, 0)
	err := GetSysApiDB(ctx, m.DB).Where("id in ?", *ids).Find(&list).Error
	return &list, err
}

func (m *SysApiRepo) Create(ctx context.Context, r *schema.Router, apiTitle string) error {
	err := GetSysApiDB(ctx, m.DB).Where(SysApi{Path: r.RelativePath, Action: r.HttpMethod}).
		Attrs(SysApi{Handle: r.Handler, Title: apiTitle}).
		FirstOrCreate(&SysApi{}).
		Error
	return err
}
