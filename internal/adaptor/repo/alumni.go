package repo

import (
	"fiber/internal/core/model"
	"fiber/internal/core/port"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type alumniRepo struct {
	db *gorm.DB
}

func NewAlumniRepo(db *gorm.DB) port.AlumniRepo {
	return &alumniRepo{
		db: db,
	}
}

func (r alumniRepo) GetById(id int) (*model.Alumni, error) {
	var result model.Alumni

	if err := r.db.
		Preload(clause.Associations).
		Take(&result, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
