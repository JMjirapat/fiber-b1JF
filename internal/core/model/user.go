package model

import (
	"time"

	"gorm.io/gorm"
)

type Flag string

const (
	FLAG_NOTFOUND Flag = "NOTFOUND"
	FLAG_FOUND    Flag = "FOUND"
	FLAG_EDIT     Flag = "EDIT"
)

type User struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	AccountID int             `gorm:"column:account_id;not null" json:"account_id"`
	Flag      Flag            `gorm:"column:flag;type:varchar(15);not null" json:"flag"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
