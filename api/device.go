package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/infrastructure"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/handler"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/repo"
	"gitlab.com/qr-through/entry/backend/internal/core/service"
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
