package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var SysUserSet = wire.NewSet(wire.Struct(new(SysUserRepo), "*"))

type SysUserRepo struct {
	DB *gorm.DB
}

func (m *SysUserRepo) GetUserByName(ctx context.Context, username string) (*SysUser, error) {
	user := new(SysUser)
	err := GetSysUserDB(ctx, m.DB).Where("username = ?", username).First(user).Error
	return user, err
}

func (m *SysUserRepo) GetPage(ctx context.Context, req *schema.SysUserGetPageReq) (*[]*SysUser, int64, error) {
	db := GetSysUserDB(ctx, m.DB).Preload("Dept").Scopes(Permission(ctx, "sys_user"))

	if req.UserID != 0 {
		db = db.Where("user_id = ?", req.UserID)
	}

	if req.Username != "" {
		db = db.Where("username = ?", req.Username)
	}

	if req.NickName != "" {
		db = db.Where("nick_name = ?", req.NickName)
	}

	if req.Phone != "" {
		db = db.Where("phone = ?", req.Phone)
	}

	if req.RoleID > 0 {
		db = db.Where("role_id = ?", req.RoleID)
	}

	if req.Sex > 0 {
		db = db.Where("sex = ?", req.Sex)
	}

	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}

	if req.PostID > 0 {
		db = db.Where("post_id = ?", req.PostID)
	}

	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}

	if req.DeptId > 0 {
		db = db.Where("dept_id = ?", req.DeptId)
	}

	if req.UserIdOrder != "" {
		db = db.Order("user_id " + req.UserIdOrder)
	}

	if req.UserIdOrder != "" {
		db = db.Order("user_id " + req.UserIdOrder)
	}

	if req.UsernameOrder != "" {
		db = db.Order("username " + req.UsernameOrder)
	}

	if req.StatusOrder != "" {
		db = db.Order("status " + req.StatusOrder)
	}

	if req.CreatedAtOrder != "" {
		db = db.Order("created_at " + req.CreatedAtOrder)
	}

	list := make([]*SysUser, 0)
	total, err := util.GetPages(db, &list, req.PageNum, req.PageSize)
	return &list, total, err
}

func (m *SysUserRepo) Get(ctx context.Context, id uint64) (*SysUser, error) {
	model := new(SysUser)
	err := GetSysUserDB(ctx, m.DB).Scopes(Permission(ctx, "sys_user")).First(model, id).Error
	return model, err
}

func (m *SysUserRepo) Insert(ctx context.Context, model *SysUser) error {
	err := GetSysUserDB(ctx, m.DB).Create(model).Error
	return err
}

func (m *SysUserRepo) Update(ctx context.Context, model *SysUser) error {
	err := GetSysUserDB(ctx, m.DB).Where("user_id = ?", model.UserID).Omit("password", "salt").Updates(model).Error
	return err
}

func (m *SysUserRepo) UpdateAvatar(ctx context.Context, id uint64, avatar string, updateBy uint64) error {
	err := GetSysUserDB(ctx, m.DB).Where("user_id =? ", id).
		Updates(map[string]interface{}{"avatar": avatar, "update_by": updateBy}).Error
	return err
}

func (m *SysUserRepo) UpdateStatus(ctx context.Context, id uint64, status int, updateBy uint64) error {
	err := GetSysUserDB(ctx, m.DB).Where("user_id =? ", id).
		Updates(map[string]interface{}{"status": status, "update_by": updateBy}).Error
	return err
}

func (m *SysUserRepo) ResetPwd(ctx context.Context, id uint64, password string, updateBy uint64) error {
	err := GetSysUserDB(ctx, m.DB).Where("user_id =? ", id).
		Updates(map[string]interface{}{"password": password, "update_by": updateBy}).Error
	return err
}

func (m *SysUserRepo) Delete(ctx context.Context, ids *[]uint64) (rowsAffected int64, err error) {
	db := GetSysUserDB(ctx, m.DB).Scopes(Permission(ctx, "sys_user")).Delete(new(SysUser), ids)
	return db.RowsAffected, db.Error
}

func (m *SysUserRepo) UpdatePwd(ctx context.Context, id uint64, password string) (rowsAffected int64, err error) {
	db := GetSysUserDB(ctx, m.DB).Where("user_id =?", id).Update("password", password)
	return db.RowsAffected, db.Error
}

func (m *SysUserRepo) FindOneWithDeptById(ctx context.Context, id uint64) (*SysUser, error) {
	model := new(SysUser)
	err := GetSysUserDB(ctx, m.DB).Preload("Dept").First(model, id).Error
	return model, err
}

func (m *SysUserRepo) DeleteSysDept(ctx context.Context, userID uint64, sysDept *[]*SysDept) error {
	err := m.DB.Model(&SysUser{UserID: userID}).Association("Dept").Delete(sysDept)
	return err
}
