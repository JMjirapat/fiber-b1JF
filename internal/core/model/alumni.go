package model

import "time"

type Alumni struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Firstname string     `gorm:"column:first_name;type:varchar(256);not null;" json:"first_name"`
	Lastname  string     `gorm:"column:last_name;type:varchar(256);not null;" json:"last_name"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
