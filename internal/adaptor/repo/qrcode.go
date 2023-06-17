package repo

import (
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
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
	var qrcodes []model.QRCodeTransaction
	if err := r.db.Preload(clause.Associations).Find(&qrcodes).Error; err != nil {
		return nil, err
	}
	return qrcodes, nil
}

func (r qrCodeRepo) GetById(id int64) (*model.QRCodeTransaction, error) {
	var qrcode model.QRCodeTransaction
	if err := r.db.
		Preload(clause.Associations).
		Take(&qrcode, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &qrcode, nil
}

func (r qrCodeRepo) UpdateById(id int64, qrcode *model.QRCodeTransaction) error {
	return r.db.Model(&model.QRCodeTransaction{}).Where("id = ?", id).Updates(&qrcode).Error
}

func (r qrCodeRepo) DeleteById(id int64) error {
	return r.db.Where("id=?", id).Delete(&model.QRCodeTransaction{}).Error
}
