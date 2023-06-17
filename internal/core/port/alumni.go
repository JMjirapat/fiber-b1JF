package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type AlumniRepo interface {
	Create(*model.Alumni) error
	GetById(id int) (*model.Alumni, error)
	UpdateById(id int, alumni model.Alumni) error
}
