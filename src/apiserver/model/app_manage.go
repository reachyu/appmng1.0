package model

import (
	"database/sql"
	"fmt"
)

type AppManageStatus string

const (
	AppCreating AppManageStatus = "CREATING"
	AppReady    AppManageStatus = "READY"
	AppDeleting AppManageStatus = "DELETING"
)

type AppManage struct {
	ID               int64          `gorm:"column:id; not null; primary_key"`
	ApplyName        string         `gorm:"column:apply_name; not null; unique"`
	ApplyType        string         `gorm:"column:apply_type; not null;"`
	ApplyFrame       string         `gorm:"column:apply_frame; not null;"`
	ApplyEnvironment string         `gorm:"column:apply_environment; not null;"`
	ApplyBrief       sql.NullString `gorm:"column:apply_brief;"`
	CreatedAt        sql.NullInt64  `gorm:"column:created_at;"`
	UpdateAt         sql.NullInt64  `gorm:"column:update_at;"`
	CreatedBy        sql.NullString `gorm:"column:created_by;"`
	UpdatedBy        sql.NullString `gorm:"column:updated_by; "`
}

func (p AppManage) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.ID)
}

func GetAppManageTablePrimaryKeyColumn() string {
	return "ID"
}

// PrimaryKeyColumnName returns the primary key for model .
func (p *AppManage) PrimaryKeyColumnName() string {
	return "ID"
}

// DefaultSortField returns the default sorting field for model .
func (p *AppManage) DefaultSortField() string {
	return "CreatedAt"
}

var AppManageAPIToModelFieldMap = map[string]string{
	"id":          "ID",
	"name":        "ApplyName",
	"created_at":  "CreatedAt",
	"description": "ApplyBrief",
}

func (p *AppManage) APIToModelFieldMap() map[string]string {
	return AppManageAPIToModelFieldMap
}
