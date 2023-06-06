package port

import "fiber/internal/core/model"

type AlumniRepo interface {
	GetById(id int) (*model.Alumni, error)
}
