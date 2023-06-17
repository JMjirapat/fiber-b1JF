package repo

import (
	"github.com/google/uuid"
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type otpRepo struct {
	db *gorm.DB
}

func NewOTPRepo(db *gorm.DB) port.OTPRepo {
	return &otpRepo{
		db: db,
	}
}

func (r otpRepo) Create(otp *model.OTPTransaction) error {
	return r.db.Create(otp).Error
}

func (r otpRepo) GetById(id uuid.UUID) (*model.OTPTransaction, error) {
	var otp model.OTPTransaction
	if err := r.db.Preload(clause.Associations).Take(&otp, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r otpRepo) All() ([]model.OTPTransaction, error) {
	var otps []model.OTPTransaction
	if err := r.db.Preload(clause.Associations).Find(&otps).Error; err != nil {
		return nil, err
	}
	return otps, nil
}

func (r otpRepo) UpdateById(id uuid.UUID, otp *model.OTPTransaction) error {
	return r.db.Where("id=?", id).Updates(&otp).Error
}

func (r otpRepo) DeleteById(id uuid.UUID) error {
	return r.db.Where("id=?", id).Delete(&model.OTPTransaction{}).Error
}
