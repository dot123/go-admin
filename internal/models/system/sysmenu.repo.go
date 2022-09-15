package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var SysMenuSet = wire.NewSet(wire.Struct(new(SysMenuRepo), "*"))

type SysMenuRepo struct {
	DB *gorm.DB
}

func (m *SysMenuRepo) GetPage(ctx context.Context, req *schema.SysMenuGetPageReq) (*[]*SysMenu, error) {
	db := GetSysMenuDB(ctx, m.DB)

	if req.Title != "" {
		db = db.Where("title = ?", req.Title)
	}

	if req.Visible > 0 {
		db = db.Where("visible = ?", req.Visible)
	}

	db = db.Order("sort ASC")
	list := make([]*SysMenu, 0)
	err := db.Preload("SysApi").Find(&list).Error

	return &list, err
}

func (m *SysMenuRepo) Get(ctx context.Context, id uint64) (*SysMenu, error) {
	model := new(SysMenu)
	err := GetSysMenuDB(ctx, m.DB).Preload("SysApi").First(model, id).Error
	return model, err
}

func (m *SysMenuRepo) Insert(ctx context.Context, model *SysMenu) error {
	err := GetSysMenuDB(ctx, m.DB).Create(model).Error
	return err
}

func (m *SysMenuRepo) UpdatePaths(ctx context.Context, menuId uint64, paths string) error {
	err := GetSysMenuDB(ctx, m.DB).Where("menu_id = ?", menuId).Update("paths", paths).Error
	return err
}

func (m *SysMenuRepo) UpdateWithAssociation(ctx context.Context, model *SysMenu) error {
	err := m.DB.Model(model).Session(&gorm.Session{FullSaveAssociations: true}).Save(model).Error
	return err
}

func (m *SysMenuRepo) DeleteSysApi(ctx context.Context, model *SysMenu) error {
	err := m.DB.Model(model).Association("SysApi").Delete(model.SysApi)
	return err
}

func (m *SysMenuRepo) FindByPath(ctx context.Context, paths string) (*[]*SysMenu, error) {
	list := make([]*SysMenu, 0)
	err := GetSysMenuDB(ctx, m.DB).Where("paths like ?", paths+"%").Find(&list).Error
	return &list, err
}

func (m *SysMenuRepo) Delete(ctx context.Context, model *[]*SysMenu) (int64, error) {
	db := m.DB.Select(clause.Associations).Delete(model)
	return db.RowsAffected, db.Error
}

func (m *SysMenuRepo) FindAll(ctx context.Context) (*[]*SysMenu, error) {
	db := GetSysMenuDB(ctx, m.DB)

	list := make([]*SysMenu, 0)
	err := db.Find(&list).Error

	return &list, err
}

func (m *SysMenuRepo) FindAllByMC(ctx context.Context) (*[]*SysMenu, error) {
	db := GetSysMenuDB(ctx, m.DB)

	list := make([]*SysMenu, 0)
	err := db.Where(" menu_type in ('M','C') and deleted_at is null").
		Order("sort").
		Find(&list).
		Error

	return &list, err
}

func (m *SysMenuRepo) FindAllByIds(ctx context.Context, ids *[]uint64) (*[]*SysMenu, error) {
	list := make([]*SysMenu, 0)
	err := GetSysMenuDB(ctx, m.DB).Where("menu_id in ?", *ids).Order("sort").Find(&list).Error
	return &list, err
}

func (m *SysMenuRepo) FindAllWithApisByIds(ctx context.Context, menuIds *[]uint64) (*[]*SysMenu, error) {
	list := make([]*SysMenu, 0)
	err := GetSysMenuDB(ctx, m.DB).Preload("SysApi").Where("menu_id in ?", *menuIds).Find(&list).Error
	return &list, err
}
