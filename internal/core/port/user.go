package port

import "fiber/internal/core/model"

type UserRepo interface {
	Create(body *model.User) error
	GetById(id int) (*model.User, error)
	All() ([]model.User, error)
	UpdateById(id int, body *model.User) error
	DeleteById(id int) error
}
