package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/internal/core/domain"
	"gitlab.com/qr-through/entry/backend/pkg/errors"
	"gitlab.com/qr-through/entry/backend/pkg/util"
)

type scannerHandler struct {
	serv domain.ScannerService
}

func NewScannerHandler(serv domain.ScannerService) *scannerHandler {
	return &scannerHandler{
		serv: serv,
	}
}

func (h scannerHandler) Verify(c *fiber.Ctx) error {
	qrCodeId := c.Params("token")
	id, err := strconv.ParseInt(qrCodeId, 10, 64)
	if err != nil {
		return util.ResponseBadRequest(c)
	}
	if err = h.serv.Verify(id); err != nil {
		return util.ResponseError(c, errors.NewUnauthorizedError(err.Error()))
	}
	return util.ResponseOK(c, nil)
}
