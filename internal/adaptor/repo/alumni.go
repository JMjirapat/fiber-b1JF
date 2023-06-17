package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
)

type alumniRepo struct {
	db *gorm.DB
}

func NewAlumniRepo(db *gorm.DB) port.AlumniRepo {
	return &alumniRepo{
		db: db,
	}
}

func (r alumniRepo) Create(alumni *model.Alumni) error {
	return r.db.Create(alumni).Error
}

func (r alumniRepo) GetById(id int) (*model.Alumni, error) {
	var alumni model.Alumni
	if err := r.db.Take(&alumni, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &alumni, nil
}

func (r alumniRepo) UpdateById(id int, alumni model.Alumni) error {
	return r.db.Where("id=?", id).Updates(&alumni).Error
}
