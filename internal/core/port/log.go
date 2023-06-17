package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type LogRepo interface {
	Create(*model.UsageLog) error
	All() ([]model.UsageLog, error)
	DeleteById(id int) error
}
