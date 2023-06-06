package repo

import (
	"fiber/internal/core/model"
	"fiber/internal/core/port"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) port.AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r accountRepo) Create(body *model.Account) (*int, error) {
	if err := r.db.Create(&body).Error; err != nil {
		return nil, err
	}
	return &body.ID, nil
}

func (r accountRepo) GetById(id int) (*model.Account, error) {
	return nil, nil
}

func (r accountRepo) GetByLineId(uid string) (*model.Account, error) {
	var result model.Account
	if err := r.db.
		Preload(clause.Associations).
		Take(&result, "line_id=?", uid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r accountRepo) All() ([]model.Account, error) {
	return nil, nil
}

func (r accountRepo) UpdateById(id int, body *model.Account) error {
	return nil
}

func (r accountRepo) DeleteById(id int) error {
	return nil
}
