package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/infrastructure"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/handler"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/repo"
	"gitlab.com/qr-through/entry/backend/internal/core/service"
)

const LINE_PREFIX = "/line"

func bindLineBotRouter(router fiber.Router) {
	lineRouter := router.Group(LINE_PREFIX)

	qrRepo := repo.NewQRCodeRepo(infrastructure.DB)
	accRepo := repo.NewAccountRepo(infrastructure.DB)
	serv := service.NewLineService(qrRepo, accRepo)
	hdl := handler.NewLineHandler(serv)

	lineRouter.Post("/webhook", hdl.Webhook)
}
