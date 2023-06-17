package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	LineID    string          `gorm:"column:line_id;type:varchar(64);not null;unique" json:"line_id"`
	Firstname string          `gorm:"column:firstname;type:varchar(256);not null;" json:"firstname"`
	Lastname  string          `gorm:"column:lastname;type:varchar(256);not null;" json:"lastname"`
	Tel       string          `gorm:"column:tel;type:varchar(15);not null;unique" json:"tel"`
	IsActive  bool            `gorm:"column:is_active;type:boolean;default:false;not null" json:"is_active"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
