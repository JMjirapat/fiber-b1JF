package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var user model.User
	if err := r.db.Preload(clause.Associations).Take(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepo) All() ([]model.User, error) {
	var users []model.User
	if err := r.db.Preload(clause.Associations).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepo) UpdateById(id int, user *model.User) error {
	return r.db.Where("id=?", id).Updates(&user).Error
}

func (r userRepo) DeleteById(id int) error {
	return r.db.Where("id=?", id).Delete(&model.User{}).Error
}
