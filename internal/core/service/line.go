package service

import (
	"errors"
	"log"
	"time"

	"gitlab.com/qr-through/entry/backend/internal/core/domain"
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
)

type lineService struct {
	qrRepo  port.QRCodeRepo
	accRepo port.AccountRepo
}

func NewLineService(qrRepo port.QRCodeRepo, accRepo port.AccountRepo) domain.LineService {
	return &lineService{
		qrRepo:  qrRepo,
		accRepo: accRepo,
	}
}

func (s lineService) CreateQR(id int64, uid string) error {
	expireTime := time.Now().Add(5 * time.Minute)

	account, err := s.accRepo.GetByLineId(uid)
	if err != nil {
		log.Panicf("%v",err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("เกิดผิดพลาด: ไม่พบข้อมูลผู้ใช้ในระบบ, กรุณาลงทะเบียนก่อน")
		}
		return errors.New("เกิดข้อผิดพลาด: ไม่สามารถสร้าง QR Code ได้ (Internal Server Error).")
	}
	qrcode := model.QRCodeTransaction{
		ID:        id,
		AccountID: account.ID,
		ExpireAt:  &expireTime,
	}

	if err = s.qrRepo.Create(&qrcode); err != nil {
		log.Panicf("%v",err)
		return errors.New("เกิดข้อผิดพลาด: ไม่สามารถสร้าง QR Code ได้ (Internal Server Error).")
	}
	return nil
}
