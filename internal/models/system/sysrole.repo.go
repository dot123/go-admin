package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var SysRoleSet = wire.NewSet(wire.Struct(new(SysRoleRepo), "*"))

type SysRoleRepo struct {
	DB *gorm.DB
}

func (m *SysRoleRepo) GetPage(ctx context.Context, req *schema.SysRoleGetPageReq) (*[]*SysRole, int64, error) {
	db := GetSysRoleDB(ctx, m.DB).Preload("SysMenu")

	if req.RoleID != 0 {
		db = db.Where("role_id = ?", req.RoleID)
	}

	if req.RoleName != "" {
		db = db.Where("role_name = ?", req.RoleName)
	}

	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}

	if req.RoleKey != "" {
		db = db.Where("role_key = ?", req.RoleKey)
	}

	if req.RoleSort > 0 {
		db = db.Where("role_sort = ?", req.RoleSort)
	}

	if req.Flag != "" {
		db = db.Where("flag = ?", req.Flag)
	}

	if req.Remark != "" {
		db = db.Where("remark = ?", req.Remark)
	}

	if req.Admin > 0 {
		db = db.Where("admin = ?", req.Admin)
	}

	if req.DataScope != "" {
		db = db.Where("data_scope = ?", req.DataScope)
	}

	list := make([]*SysRole, 0)
	total, err := util.GetPages(db, &list, req.PageNum, req.PageSize)
	return &list, total, err
}

func (m *SysRoleRepo) Get(ctx context.Context, id uint64) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).First(model, id).Error
	return model, err
}

func (m *SysRoleRepo) FindOneWithSysMenu(ctx context.Context, id uint64) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Preload("SysMenu").First(model, id).Error
	return model, err
}

func (m *SysRoleRepo) FindOneWithSysMenuByRoleKey(ctx context.Context, roleKey string) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Where("role_key = ? ", roleKey).Preload("SysMenu").First(model).Error
	return model, err
}

func (m *SysRoleRepo) FindOneByRoleKey(ctx context.Context, roleKey string) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Where("role_key = ? ", roleKey).First(model).Error
	return model, err
}

func (m *SysRoleRepo) FindOneByRoleName(ctx context.Context, roleName string) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Where("role_name = ? ", roleName).First(model).Error
	return model, err
}

func (m *SysRoleRepo) Insert(ctx context.Context, model *SysRole) error {
	err := GetSysRoleDB(ctx, m.DB).Save(model).Error
	return err
}

func (m *SysRoleRepo) DeleteSysMenu(ctx context.Context, roleID uint64, menus *[]*SysMenu) error {
	err := m.DB.Model(&SysRole{RoleID: roleID}).Association("SysMenu").Delete(menus)
	return err
}

func (m *SysRoleRepo) UpdateWithAssociation(ctx context.Context, model *SysRole) error {
	err := m.DB.Model(model).Session(&gorm.Session{FullSaveAssociations: true}).Save(model).Error
	return err
}

func (m *SysRoleRepo) UpdateStatus(ctx context.Context, id uint64, status int, updateBy uint64) error {
	err := GetSysRoleDB(ctx, m.DB).Where("role_id =? ", id).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(map[string]interface{}{"status": status, "update_by": updateBy}).Error
	return err
}

func (m *SysRoleRepo) FindOneWithSysMenuAndSysDept(ctx context.Context, id uint64) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Preload("SysMenu").Preload("SysDept").First(model, id).Error
	return model, err
}

func (m *SysRoleRepo) Delete(ctx context.Context, model *SysRole) (int64, error) {
	db := m.DB.Select(clause.Associations).Delete(model)
	return db.RowsAffected, db.Error
}

func (m *SysRoleRepo) FindOneWithSysDept(ctx context.Context, id uint64) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Preload("SysDept").First(model, id).Error
	return model, err
}

func (m *SysRoleRepo) DeleteSysDept(ctx context.Context, roleID uint64, sysDept *[]*SysDept) error {
	err := m.DB.Model(&SysRole{RoleID: roleID}).Association("SysDept").Delete(sysDept)
	return err
}

func (m *SysRoleRepo) FindAllBy(ctx context.Context, roleName string) (*SysRole, error) {
	model := new(SysRole)
	err := GetSysRoleDB(ctx, m.DB).Where("role_name = ? ", roleName).First(model).Error
	return model, err
}
