package handler

import (
	"fiber/internal/core/domain"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
		log.Panic(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err = h.serv.Verify(id); err != nil {
		return c.Status(401).SendString(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
