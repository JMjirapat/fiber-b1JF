package model

import (
	"time"

	"gorm.io/gorm"
)

type UsageLog struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   int             `gorm:"column:account_id;not null" json:"account_id"`
	AccountInfo string          `gorm:"column:account_info;type:varchar(15);not null" json:"account_info"`
	CreatedAt   *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
