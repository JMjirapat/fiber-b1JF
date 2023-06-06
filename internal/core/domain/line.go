package domain

type LineService interface {
	CreateQR(id int64, uid string) error
	RegisterUser(body LineRegisterBody, lineID string) error
	GetById(id int) (*LIFFGetByIdResponse, error)
}

type LIFFGetByIdResponse struct {
	InAlumni    bool    `json:"in_alumni"`
	StudentCode string  `json:"student_code"`
	Firstname   *string `json:"first_name"`
	Lastname    *string `json:"last_name"`
}

type LineRegisterBody struct {
	StudentCode string `json:"student_code"`
	TokenID     string `json:"token_id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Tel         string `json:"tel"`
}
