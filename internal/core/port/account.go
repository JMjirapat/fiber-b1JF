package port

import "fiber/internal/core/model"

type AccountRepo interface {
	Create(body *model.Account) (*int, error)
	GetById(id int) (*model.Account, error)
	GetByLineId(uid string) (*model.Account, error)
	All() ([]model.Account, error)
	UpdateById(id int, body *model.Account) error
	DeleteById(id int) error
}
