package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
)

type alumniNewRepo struct {
	db *gorm.DB
}

func NewAlumniNewRepo(db *gorm.DB) port.AlumniNewRepo {
	return &alumniNewRepo{
		db: db,
	}
}

func (r alumniNewRepo) Create(alumni *model.Alumni_new) error {
	return r.db.Create(alumni).Error
}
