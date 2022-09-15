package system

import (
	"context"
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"gorm.io/gorm"
)

func Permission(ctx context.Context, tableName string) func(db *gorm.DB) *gorm.DB {
	p := contextx.FromDataPermission(ctx)

	return func(db *gorm.DB) *gorm.DB {
		switch p.DataScope {
		case "2":
			return db.Where(tableName+".create_by in (select sys_user.user_id from sys_role_dept left join sys_user on sys_user.dept_id=sys_role_dept.dept_id where sys_role_dept.role_id = ?)", p.RoleID)
		case "3":
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where dept_id = ? )", p.DeptID)
		case "4":
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where sys_user.dept_id in(select dept_id from sys_dept where dept_path like ? ))", "%/"+fmt.Sprintf("%d", p.DeptID)+"/%")
		case "5":
			return db.Where(tableName+".create_by = ?", p.UserID)
		default:
			return db
		}
	}
}
