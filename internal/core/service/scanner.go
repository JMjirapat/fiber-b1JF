package service

import (
	"errors"
	"fiber/internal/core/domain"
	"fiber/internal/core/model"
	"fiber/internal/core/port"
	"log"
	"time"
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
		return errors.New("not found QRCode")
	}
	if result.IsUsed {
		return errors.New("QRCode Already Used")
	}

	if (*result.ExpireAt).Before(time.Now()) {
		return errors.New("QR Code Expired")
	}
	data := model.QRCodeTransaction{
		IsUsed: true,
	}
	if err = s.qrRepo.UpdateById(id, &data); err != nil {
		return errors.New("server error")
	}

	usage := model.UsageLog{
		AccountID: result.AccountID,
	}

	if err = s.logRepo.Create(&usage); err != nil {
		log.Panic(err)
		return errors.New("Something went Wrong")
	}
	return nil
}
