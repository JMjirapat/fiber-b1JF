package api

import (
	"fiber/internal/adaptor/handler"
	"fiber/internal/adaptor/repo"
	"fiber/internal/core/service"

	"fiber/infrastructure"

	"github.com/gofiber/fiber/v2"
)

const LINE_PREFIX = "/line"

func bindLineBotRouter(router fiber.Router) {
	lineRouter := router.Group(LINE_PREFIX)

	qrRepo := repo.NewQRCodeRepo(infrastructure.DB)
	accRepo := repo.NewAccountRepo(infrastructure.DB)
	userRepo := repo.NewUserRepo(infrastructure.DB)
	alumniRepo := repo.NewAlumniRepo(infrastructure.DB)
	serv := service.NewLineService(qrRepo, accRepo, userRepo, alumniRepo)
	hdl := handler.NewLineHandler(serv)

	lineRouter.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
		})
	})
	lineRouter.Post("/webhook", hdl.Webhook)
	lineRouter.Get("/user/check/:id", hdl.GetById)
	lineRouter.Post("/user/register", hdl.Register)
}
