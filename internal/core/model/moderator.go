package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	ROLE_MODERATOR Role = "MODERATOR"
	ROLE_ADMIN     Role = "ADMIN"
)

type Moderator struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID int             `gorm:"column:account_id;not null" json:"account_id"`
	Role      Role            `gorm:"column:role;type:varchar(9);not null;" json:"role"`
	IsActive  bool            `gorm:"column:is_active;type:boolean;default:true;not null" json:"is_active"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
