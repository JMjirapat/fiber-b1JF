package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type UserRepo interface {
	Create(*model.User) error
	GetById(id int) (*model.User, error)
	All() ([]model.User, error)
	UpdateById(id int, user *model.User) error
	DeleteById(id int) error
}
