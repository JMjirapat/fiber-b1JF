package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type moderatorRepo struct {
	db *gorm.DB
}

func NewModeratorRepo(db *gorm.DB) port.ModeratorRepo {
	return &moderatorRepo{
		db: db,
	}
}

func (r moderatorRepo) Create(moderator *model.Moderator) error {
	return r.db.Create(moderator).Error
}

func (r moderatorRepo) GetById(id int) (*model.Moderator, error) {
	var moderator model.Moderator
	if err := r.db.Preload(clause.Associations).Take(&moderator, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &moderator, nil
}

func (r moderatorRepo) All() ([]model.Moderator, error) {
	var moderators []model.Moderator
	if err := r.db.Preload(clause.Associations).Find(&moderators).Error; err != nil {
		return nil, err
	}
	return moderators, nil
}

func (r moderatorRepo) UpdateById(id int, moderator *model.Moderator) error {
	return r.db.Where("id=?", id).Updates(&moderator).Error
}

func (r moderatorRepo) DeleteById(id int) error {
	return r.db.Where("id=?", id).Delete(&model.Moderator{}).Error
}
