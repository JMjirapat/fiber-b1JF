package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
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
	if err := r.db.Create(body).Error; err != nil {
		return nil, err
	}
	return &body.ID, nil
}

func (r accountRepo) GetById(id int) (*model.Account, error) {
	var account model.Account
	if err := r.db.Take(&account, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r accountRepo) GetByLineId(uid string) (*model.Account, error) {
	var account model.Account
	if err := r.db.Take(&account, "line_id=?", uid).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (r accountRepo) All() ([]model.Account, error) {
	var accounts []model.Account
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r accountRepo) UpdateById(id int, account *model.Account) error {
	return r.db.Where("id=?", id).Updates(&account).Error
}

func (r accountRepo) DeleteById(id int) error {
	return r.db.Where("id=?", id).Delete(&model.Account{}).Error
}
