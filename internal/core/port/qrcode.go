package port

import "fiber/internal/core/model"

type QRCodeRepo interface {
	Create(body *model.QRCodeTransaction) error
	GetById(id int64) (*model.QRCodeTransaction, error)
	All() ([]model.QRCodeTransaction, error)
	UpdateById(id int64, body *model.QRCodeTransaction) error
	DeleteById(id int64) error
}
