package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/dot123/gin-gorm-admin/internal/config"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) (*casbin.SyncedEnforcer, func(), error) {
	cfg := config.C.Casbin

	e, err := NewCasbin(db, cfg.Debug, cfg.Model)
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	return e, cleanFunc, nil
}

func NewCasbin(db *gorm.DB, debug bool, modelPath string) (*casbin.SyncedEnforcer, error) {
	type CasbinRule struct {
		ID    uint   `gorm:"primaryKey;autoIncrement"`
		Ptype string `gorm:"size:512;"`
		V0    string `gorm:"size:512;"`
		V1    string `gorm:"size:512;"`
		V2    string `gorm:"size:512;"`
		V3    string `gorm:"size:512;"`
		V4    string `gorm:"size:512;"`
		V5    string `gorm:"size:512;"`
	}

	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, new(CasbinRule))
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewSyncedEnforcer(modelPath, a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}

	e.EnableLog(debug)

	return e, nil
}
