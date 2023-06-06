package api

import (
	"fiber/internal/adaptor/handler"
	"fiber/internal/adaptor/repo"
	"fiber/internal/core/service"

	"fiber/infrastructure"

	"github.com/gofiber/fiber/v2"
)

const DEVICE_PREFIX = "/device"

func bindDeviceRouter(router fiber.Router) {
	deviceRouter := router.Group(DEVICE_PREFIX)

	qrRepo := repo.NewQRCodeRepo(infrastructure.DB)
	logRepo := repo.NewLogRepo(infrastructure.DB)
	serv := service.NewScannerService(qrRepo, logRepo)
	hdl := handler.NewScannerHandler(serv)

	deviceRouter.Get("/scanner/:token", hdl.Verify)
}
