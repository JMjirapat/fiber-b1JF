package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) port.LogRepo {
	return &logRepo{
		db: db,
	}
}

func (r logRepo) Create(log *model.UsageLog) error {
	return r.db.Create(log).Error
}

func (r logRepo) All() ([]model.UsageLog, error) {
	var logs []model.UsageLog
	if err := r.db.Preload(clause.Associations).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r logRepo) DeleteById(id int) error {
	return r.db.Where("id=?", id).Delete(&model.UsageLog{}).Error
}
