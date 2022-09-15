package contextx

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/schema"
)

type (
	transCtx          struct{}
	noTransCtx        struct{}
	transLockCtx      struct{}
	userIDCtx         struct{}
	userNameCtx       struct{}
	roleKeyCtx        struct{}
	dataPermissionCtx struct{}
	traceIDCtx        struct{}
)

// Wrap transaction context
func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

func NewNoTrans(ctx context.Context) context.Context {
	return context.WithValue(ctx, noTransCtx{}, true)
}

func FromNoTrans(ctx context.Context) bool {
	v := ctx.Value(noTransCtx{})
	return v != nil && v.(bool)
}

func NewTransLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, true)
}

func FromTransLock(ctx context.Context) bool {
	v := ctx.Value(transLockCtx{})
	return v != nil && v.(bool)
}

func NewUserID(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, userIDCtx{}, userID)
}

func FromUserID(ctx context.Context) uint64 {
	v := ctx.Value(userIDCtx{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}

func NewUserName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameCtx{}, userName)
}

func FromUserName(ctx context.Context) string {
	v := ctx.Value(userNameCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewRoleKey(ctx context.Context, roleKey string) context.Context {
	return context.WithValue(ctx, roleKeyCtx{}, roleKey)
}

func FromRoleKey(ctx context.Context) string {
	v := ctx.Value(roleKeyCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewDataPermission(ctx context.Context, p *schema.DataPermission) context.Context {
	return context.WithValue(ctx, dataPermissionCtx{}, p)
}

func FromDataPermission(ctx context.Context) *schema.DataPermission {
	v := ctx.Value(dataPermissionCtx{})
	if v != nil {
		if s, ok := v.(*schema.DataPermission); ok {
			return s
		}
	}
	return nil
}

func NewTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtx{}, traceID)
}

func FromTraceID(ctx context.Context) (string, bool) {
	v := ctx.Value(traceIDCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s, s != ""
		}
	}
	return "", false
}
