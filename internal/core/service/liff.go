package service

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/qr-through/entry/backend/internal/core/domain"
	"gitlab.com/qr-through/entry/backend/internal/core/model"
	"gitlab.com/qr-through/entry/backend/internal/core/port"
	"gorm.io/gorm"
)

type liffService struct {
	accRepo       port.AccountRepo
	userRepo      port.UserRepo
	alumniRepo    port.AlumniRepo
	alumniNewRepo port.AlumniNewRepo
	otpRepo       port.OTPRepo
}

func NewLiffService(accRepo port.AccountRepo, userRepo port.UserRepo, alumniRepo port.AlumniRepo, alumniNewRepo port.AlumniNewRepo, otpRepo port.OTPRepo) domain.LiffService {
	return &liffService{
		accRepo:       accRepo,
		userRepo:      userRepo,
		alumniRepo:    alumniRepo,
		alumniNewRepo: alumniNewRepo,
		otpRepo:       otpRepo,
	}
}

func (s liffService) GetAlumniById(id int) (*domain.LIFFGetAlumniByIdResponse, error) {
	result, err := s.alumniRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			alumni := domain.LIFFGetAlumniByIdResponse{
				InAlumni:    false,
				StudentCode: strconv.Itoa(id),
			}
			return &alumni, nil
		}
		log.Panic(err)
		return nil, errors.New("ไม่สามารถตรวจสอบรหัสนักศีกษาได้")
	}

	alumni := domain.LIFFGetAlumniByIdResponse{
		InAlumni:    true,
		StudentCode: strconv.Itoa(id),
		Firstname:   &result.Firstname,
		Lastname:    &result.Lastname,
		Tel:         &result.Tel,
	}

	return &alumni, nil

}

func (s liffService) SignUp(body domain.LIFFRegisterBody, lineID string) error {
	account := model.Account{
		LineID:    lineID,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Tel:       body.Tel,
	}

	code, err := strconv.Atoi(body.StudentCode)
	if err != nil {
		return errors.New("เกิดข้อผิดพลาด ไม่สามารถลงทะเบียนใช้งานได้1")
	}

	acc, err := s.accRepo.Create(&account)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_line_id_key\"") {
			return errors.New("line นี้เคยลงทะเบียนในระบบแล้ว")
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_tel_key\"") {
			return errors.New("เบอร์โทรศัพท์นี้เคยลงทะเบียนในระบบแล้ว")
		}
		log.Panic(err)
		return errors.New("เกิดข้อผิดพลาด ไม่สามารถลงทะเบียนใช้งานได้2")
	}

	var flag = model.FLAG_NOTFOUND

	alumni, err := s.alumniRepo.GetById(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			flag = model.FLAG_NOTFOUND
			if err = s.alumniNewRepo.Create(&model.Alumni_new{
				ID:        code,
				Firstname: body.Firstname,
				Lastname:  body.Lastname,
				Tel:       body.Tel,
			}); err != nil {
				log.Panic(err)
			}
		}
	} else {
		flag = model.FLAG_FOUND
		if body.Firstname != alumni.Firstname || body.Lastname != alumni.Lastname || body.Tel != alumni.Tel {
			flag = model.FLAG_EDIT
			if err = s.alumniNewRepo.Create(&model.Alumni_new{
				ID:        code,
				Firstname: body.Firstname,
				Lastname:  body.Lastname,
				Tel:       body.Tel,
			}); err != nil {
				log.Panic(err)
			}
		}
	}

	user := model.User{
		ID:        code,
		AccountID: *acc,
		Flag:      flag,
	}
	if err = s.userRepo.Create(&user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("รหัสนักศึกษานี้เคยลงทะเบียนในระบบแล้ว")
		}
		log.Panic(err)
		return errors.New("เกิดข้อผิดพลาด ไม่สามารถลงทะเบียนใช้งานได้3")
	}
	return nil
}

func (s liffService) GetOTP(body domain.LIFFGetOTPBody) (*domain.LIFFGetOTPResponse, error) {
	// Thrid Party OTP Service | Request OTP
	id := uuid.New()
	otp := model.OTPTransaction{
		ID:  id,
		Tel: body.Tel,
	}

	if err := s.otpRepo.Create(&otp); err != nil {
		log.Panic(err)
		//อนาคตเอาออก
		return nil, errors.New("ไม่สามารถสร้าง OTP Transaction ได้")
	}

	res := domain.LIFFGetOTPResponse{
		Token: id,
		Refno: strconv.Itoa(rand.Intn(9999)),
	}

	return &res, nil
}

func (s liffService) VerifyOTP(body domain.LIFFVerifyOTPBody) (*string, error) {
	// Thrid Party OTP Service | Verify OTP
	result, err := s.otpRepo.GetById(body.Token)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("ไม่เจอ OTP นี้ในระบบ, กรุณาขอ OTP ใหม่ภายหลัง")
		}
		log.Panic(err)
		return nil, errors.New("ไม่สามารถตรวจสอบ OTP ได้")
	}

	if result.IsUsed {
		return nil, errors.New("OTP token นี้เคยใช้งานแล้ว")
	}

	if body.Pin != "4444" {
		return nil, errors.New("รหัส OTP ไม่ถูกต้องกรุณากรอกใหม่")
	}

	otp := model.OTPTransaction{
		IsUsed: true,
	}

	if err := s.otpRepo.UpdateById(body.Token, &otp); err != nil {
		log.Panic(err)
		//return nil, errors.New("ไม่สามารถอัพเดต OTP ในระบบได้, กรุณาติดต่อเจ้าหน้าที่")
	}

	res := "success"

	return &res, nil
}
