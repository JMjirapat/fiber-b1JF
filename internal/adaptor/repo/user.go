package repo

import (
	"fiber/internal/core/model"
	"fiber/internal/core/port"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(body *model.User) error {
	return r.db.Create(body).Error
}

func (r userRepo) GetById(id int) (*model.User, error) {
	return nil, nil
}

func (r userRepo) All() ([]model.User, error) {
	return nil, nil
}

func (r userRepo) UpdateById(id int, body *model.User) error {
	return nil
}

func (r userRepo) DeleteById(id int) error {
	return nil
}
