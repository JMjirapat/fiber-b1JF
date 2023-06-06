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

type Account struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	LineID    string          `gorm:"column:line_id;type:varchar(64);not null;unique" json:"line_id"`
	Firstname string          `gorm:"column:first_name;type:varchar(256);not null;" json:"first_name"`
	Lastname  string          `gorm:"column:last_name;type:varchar(256);not null;" json:"last_name"`
	Tel       string          `gorm:"column:tel;type:varchar(10);not null;" json:"tel"`
	IsActive  bool            `gorm:"column:is_active;type:boolean;default:false;not null" json:"is_active"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type User struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	AccountID int             `gorm:"column:account_id;not null" json:"account_id"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}

type Moderator struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID int             `gorm:"column:account_id;not null" json:"account_id"`
	Role      Role            `gorm:"column:role;type:varchar(20);not null;" json:"role"`
	IsActive  bool            `gorm:"column:is_active;type:boolean;default:true;not null" json:"is_active"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
