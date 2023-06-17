package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type AlumniNewRepo interface {
	Create(*model.Alumni_new) error
}
