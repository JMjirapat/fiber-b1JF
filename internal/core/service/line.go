package service

import (
	"errors"
	"fiber/internal/core/domain"
	"fiber/internal/core/model"
	"fiber/internal/core/port"
	"log"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type lineService struct {
	qrRepo     port.QRCodeRepo
	accRepo    port.AccountRepo
	userRepo   port.UserRepo
	alumniRepo port.AlumniRepo
}

func NewLineService(qrRepo port.QRCodeRepo, accRepo port.AccountRepo, userRepo port.UserRepo, alumniRepo port.AlumniRepo) domain.LineService {
	return &lineService{
		qrRepo:     qrRepo,
		accRepo:    accRepo,
		userRepo:   userRepo,
		alumniRepo: alumniRepo,
	}
}

func (s lineService) CreateQR(id int64, uid string) error {
	expireTime := time.Now().Add(1 * time.Minute)

	account, err := s.accRepo.GetByLineId(uid)
	if err != nil {
		return errors.New("not found line userid, please register first")
	}
	qrcode := model.QRCodeTransaction{
		ID:        id,
		AccountID: account.ID,
		ExpireAt:  &expireTime,
	}

	if err = s.qrRepo.Create(&qrcode); err != nil {
		log.Panic(err)
		return errors.New("Can't Generate QRCode.")
	}
	return nil
}

func (s lineService) GetById(id int) (*domain.LIFFGetByIdResponse, error) {
	result, err := s.alumniRepo.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			alumni := domain.LIFFGetByIdResponse{
				InAlumni:    false,
				StudentCode: strconv.Itoa(id),
			}
			return &alumni, nil
		}
		return nil, errors.New("can't get student dode")
	}

	alumni := domain.LIFFGetByIdResponse{
		InAlumni:    true,
		StudentCode: strconv.Itoa(result.ID),
		Firstname:   &result.Firstname,
		Lastname:    &result.Lastname,
	}

	return &alumni, nil

}

func (s lineService) RegisterUser(body domain.LineRegisterBody, lineID string) error {
	account := model.Account{
		LineID:    lineID,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Tel:       body.Tel,
	}

	acc, err := s.accRepo.Create(&account)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_line_id_key\"") {
			return errors.New("line userid already registered")
		}
		return errors.New("something went Wrong")
	}

	code, err := strconv.Atoi(body.StudentCode)
	if err != nil {
		return errors.New("something went Wrong")
	}

	user := model.User{
		ID:        code,
		AccountID: *acc,
	}
	if err = s.userRepo.Create(&user); err != nil {
		log.Panic(err)
		if err == gorm.ErrDuplicatedKey {
			return errors.New("Duplicate User Key.")
		}
		return errors.New("Something went Wrong")
	}
	return nil
}
