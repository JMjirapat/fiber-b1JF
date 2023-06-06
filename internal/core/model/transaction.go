package model

import (
	"time"

	"gorm.io/gorm"
)

type QRCodeTransaction struct {
	ID        int64           `gorm:"primaryKey" json:"id"`
	AccountID int             `gorm:"column:account_id;not null" json:"account_id"`
	IsUsed    bool            `gorm:"column:is_used;type:boolean;default:false;not null" json:"is_used"`
	ExpireAt  *time.Time      `gorm:"column:expire_at;not null" json:"expire_at"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}

type OTPTranstaction struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Tel       string          `gorm:"column:tel;type:varchar(10);not null;" json:"tel"`
	IsUsed    bool            `gorm:"column:is_used;type:boolean;default:false;not null" json:"is_used"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
