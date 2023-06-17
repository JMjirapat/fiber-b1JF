package port

import (
	"github.com/google/uuid"
	"gitlab.com/qr-through/entry/backend/internal/core/model"
)

type OTPRepo interface {
	Create(*model.OTPTransaction) error
	GetById(id uuid.UUID) (*model.OTPTransaction, error)
	All() ([]model.OTPTransaction, error)
	UpdateById(id uuid.UUID, otp *model.OTPTransaction) error
	DeleteById(id uuid.UUID) error
}
