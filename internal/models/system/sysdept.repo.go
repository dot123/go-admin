package system

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var SysDeptSet = wire.NewSet(wire.Struct(new(SysDeptRepo), "*"))

type SysDeptRepo struct {
	DB *gorm.DB
}

func (m *SysDeptRepo) GetPage(ctx context.Context, req *schema.SysDeptGetPageReq) (*[]*SysDept, error) {
	db := GetSysDeptDB(ctx, m.DB)

	if req.DeptID != 0 {
		db = db.Where("dept_id = ?", req.DeptID)
	}

	if req.ParentID > 0 {
		db = db.Where("parent_id = ?", req.ParentID)
	}

	if req.DeptPath != "" {
		db = db.Where("dept_path = ?", req.DeptPath)
	}

	if req.DeptName != "" {
		db = db.Where("dept_name = ?", req.DeptName)
	}

	if req.Sort > 0 {
		db = db.Where("sort = ?", req.Sort)
	}

	if req.Leader != "" {
		db = db.Where("leader = ?", req.Leader)
	}

	if req.Phone != "" {
		db = db.Where("phone = ?", req.Phone)
	}

	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}

	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}

	list := make([]*SysDept, 0)
	err := db.Find(&list).Error
	return &list, err
}

func (m *SysDeptRepo) Get(ctx context.Context, id uint64) (*SysDept, error) {
	model := new(SysDept)
	err := GetSysDeptDB(ctx, m.DB).First(model, id).Error
	return model, err
}

func (m *SysDeptRepo) Insert(ctx context.Context, model *SysDept) error {
	err := GetSysDeptDB(ctx, m.DB).Create(model).Error
	return err
}

func (m *SysDeptRepo) UpdateDeptPath(ctx context.Context, id uint64, deptPath string) error {
	err := GetSysDeptDB(ctx, m.DB).Where("dept_id = ?", id).Update("dept_path", deptPath).Error
	return err
}

func (m *SysDeptRepo) Update(ctx context.Context, model *SysDept) error {
	err := m.DB.Save(model).Error
	return err
}

func (m *SysDeptRepo) Delete(ctx context.Context, model *[]*SysDept) (rowsAffected int64, err error) {
	db := m.DB.Select(clause.Associations).Delete(model)
	return db.RowsAffected, db.Error
}

func (m *SysDeptRepo) FindAll(ctx context.Context) (*[]*SysDept, error) {
	list := make([]*SysDept, 0)
	err := GetSysDeptDB(ctx, m.DB).Find(&list).Error
	return &list, err
}

func (m *SysDeptRepo) GetWithRoleId(ctx context.Context, roleId uint64) (*[]*schema.DeptIDList, error) {
	deptList := make([]*schema.DeptIDList, 0)
	if err := m.DB.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, err
	}

	return &deptList, nil
}

func (m *SysDeptRepo) FindAllByIds(ctx context.Context, ids *[]uint64) (*[]*SysDept, error) {
	list := make([]*SysDept, 0)
	err := GetSysDeptDB(ctx, m.DB).Where("dept_id in ?", *ids).Find(&list).Error
	return &list, err
}
