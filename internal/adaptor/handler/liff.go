package handler

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/config"
	"gitlab.com/qr-through/entry/backend/internal/core/domain"
	"gitlab.com/qr-through/entry/backend/pkg/errors"
	"gitlab.com/qr-through/entry/backend/pkg/util"
)

type Token struct {
	Iss string   `json:"iss"`
	Sub string   `json:"sub"`
	Aud string   `json:"aud"`
	Exp int64    `json:"exp"`
	Iat int64    `json:"iat"`
	Amr []string `json:"amr"`
}

type liffHandler struct {
	serv domain.LiffService
}

func NewLiffHandler(serv domain.LiffService) *liffHandler {
	return &liffHandler{
		serv: serv,
	}
}

func (h liffHandler) GetAlumniById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.ResponseBadRequest(c)
	}

	result, err := h.serv.GetAlumniById(id)
	if err != nil {
		return util.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return util.ResponseOK(c, &result)
}

func (h liffHandler) SignUp(c *fiber.Ctx) error {
	cfg := config.Config

	var body domain.LIFFRegisterBody
	if err := c.BodyParser(&body); err != nil {
		return util.ResponseUnprocessableEntity(c)
	}

	formData := url.Values{}
	formData.Set("id_token", body.TokenID)
	formData.Set("client_id", cfg.ChannelID)

	res, err := util.HttpPOST[Token]("https://api.line.me/oauth2/v2.1/verify", formData.Encode())
	if err != nil {
		return util.ResponseError(c, errors.NewStatusBadGatewayError("เกิดข้อผิดพลาดกับการติดต่อบริการของ Line"))
	}

	if err := h.serv.SignUp(body, res.Sub); err != nil {
		return util.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return util.ResponseCreated(c, nil)
}

func (h liffHandler) GetOTP(c *fiber.Ctx) error {
	var body domain.LIFFGetOTPBody
	if err := c.BodyParser(&body); err != nil {
		return util.ResponseUnprocessableEntity(c)
	}

	result, err := h.serv.GetOTP(body)
	if err != nil {
		return util.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return util.ResponseCreated(c, result)
}

func (h liffHandler) VerifyOTP(c *fiber.Ctx) error {
	var body domain.LIFFVerifyOTPBody
	if err := c.BodyParser(&body); err != nil {
		return util.ResponseUnprocessableEntity(c)
	}

	result, err := h.serv.VerifyOTP(body)
	if err != nil {
		return util.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return util.Response(c, fiber.StatusCreated, *result, nil)
}
