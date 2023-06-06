package repo

import (
	"fiber/internal/core/model"
	"fiber/internal/core/port"

	"gorm.io/gorm"
)

type logRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) port.LogRepo {
	return &logRepo{
		db: db,
	}
}

func (r logRepo) Create(body *model.UsageLog) error {
	return r.db.Create(body).Error
}
