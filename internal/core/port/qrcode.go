package port

import "gitlab.com/qr-through/entry/backend/internal/core/model"

type QRCodeRepo interface {
	Create(*model.QRCodeTransaction) error
	GetById(id int64) (*model.QRCodeTransaction, error)
	All() ([]model.QRCodeTransaction, error)
	UpdateById(id int64, qrcode *model.QRCodeTransaction) error
	DeleteById(id int64) error
}
