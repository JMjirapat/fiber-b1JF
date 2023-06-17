package domain

import "github.com/google/uuid"

type LiffService interface {
	SignUp(body LIFFRegisterBody, lineID string) error
	GetAlumniById(id int) (*LIFFGetAlumniByIdResponse, error)
	GetOTP(body LIFFGetOTPBody) (*LIFFGetOTPResponse, error)
	VerifyOTP(body LIFFVerifyOTPBody) (*string, error)
}

type LIFFGetAlumniByIdResponse struct {
	InAlumni    bool    `json:"in_alumni"`
	StudentCode string  `json:"student_code"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	Tel         *string `json:"tel"`
}

type LIFFRegisterBody struct {
	StudentCode string `json:"student_code"`
	TokenID     string `json:"token_id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Tel         string `json:"tel"`
}

type LIFFGetOTPBody struct {
	Tel string `json:"tel"`
}

type LIFFGetOTPResponse struct {
	Refno string    `json:"refno"`
	Token uuid.UUID `json:"token"`
}

type LIFFVerifyOTPBody struct {
	Token uuid.UUID `json:"token"`
	Pin   string    `json:"pin"`
}
