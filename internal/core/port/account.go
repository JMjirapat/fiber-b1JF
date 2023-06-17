package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type AccountRepo interface {
	Create(account *model.Account) (*int, error)
	GetById(id int) (*model.Account, error)
	GetByLineId(uid string) (*model.Account, error)
	All() ([]model.Account, error)
	UpdateById(id int, account *model.Account) error
	DeleteById(id int) error
}
