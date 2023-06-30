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

type scannerService struct {
	qrRepo  port.QRCodeRepo
	logRepo port.LogRepo
}

func NewScannerService(qrRepo port.QRCodeRepo, logRepo port.LogRepo) domain.ScannerService {
	return &scannerService{
		qrRepo:  qrRepo,
		logRepo: logRepo,
	}
}

func (s scannerService) Verify(id int64) error {
	result, err := s.qrRepo.GetById(id)
	if err != nil {
		log.Printf("%v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("qrcode not found")
		}
		return err
	}

	if (*result.ExpireAt).Before(time.Now().UTC().Add(time.Hour * 7)) {
		log.Printf("%v", *result.ExpireAt)
		log.Printf("%v", (*result.ExpireAt).Before(time.Now().UTC().Add(time.Hour*7)))
		log.Printf("%v", time.Now().UTC().Add(time.Hour*7))
		log.Printf("%v", time.Now().Add(time.Hour*7))
		log.Printf("%v", time.Now())
		return errors.New("qrcode expired")
	}

	data := model.QRCodeTransaction{
		NumUsed: result.NumUsed + 1,
	}
	if err = s.qrRepo.UpdateById(id, &data); err != nil {
		return err
	}

	usage := model.UsageLog{
		AccountID: result.AccountID,
	}
	if err = s.logRepo.Create(&usage); err != nil {
		log.Panic(err)
		return errors.New("can't create usage log.")
	}
	return nil
}
