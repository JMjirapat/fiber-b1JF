package repo

import (
	"fiber/internal/core/model"
	"fiber/internal/core/port"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type qrCodeRepo struct {
	db *gorm.DB
}

func NewQRCodeRepo(db *gorm.DB) port.QRCodeRepo {
	return &qrCodeRepo{
		db: db,
	}
}

func (r qrCodeRepo) Create(body *model.QRCodeTransaction) error {
	return r.db.Create(body).Error
}

func (r qrCodeRepo) All() ([]model.QRCodeTransaction, error) {
	return nil, nil
}

func (r qrCodeRepo) GetById(id int64) (*model.QRCodeTransaction, error) {
	var result model.QRCodeTransaction
	if err := r.db.
		Preload(clause.Associations).
		Take(&result, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r qrCodeRepo) UpdateById(id int64, body *model.QRCodeTransaction) error {
	if err := r.db.Model(&model.QRCodeTransaction{}).Where("id = ?", id).Omit("id").Updates(body).Error; err != nil {
		return err
	}
	return nil
}

func (r qrCodeRepo) DeleteById(id int64) error {
	return nil
}
