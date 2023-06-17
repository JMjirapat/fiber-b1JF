package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type ModeratorRepo interface {
	Create(moderator *model.Moderator) error
	GetById(id int) (*model.Moderator, error)
	All() ([]model.Moderator, error)
	UpdateById(id int, moderator *model.Moderator) error
	DeleteById(id int) error
}
