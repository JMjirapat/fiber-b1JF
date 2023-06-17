package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/infrastructure"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/handler"
	"gitlab.com/qr-through/entry/backend/internal/adaptor/repo"
	"gitlab.com/qr-through/entry/backend/internal/core/service"
)

const LIFF_PREFIX = "/liff"

func bindLiffRouter(router fiber.Router) {
	liffRouter := router.Group(LIFF_PREFIX)

	accRepo := repo.NewAccountRepo(infrastructure.DB)
	userRepo := repo.NewUserRepo(infrastructure.DB)
	alumniRepo := repo.NewAlumniRepo(infrastructure.DB)
	alumniNewRepo := repo.NewAlumniNewRepo(infrastructure.DB)
	otpRepo := repo.NewOTPRepo(infrastructure.DB)
	serv := service.NewLiffService(accRepo, userRepo, alumniRepo, alumniNewRepo, otpRepo)
	hdl := handler.NewLiffHandler(serv)

	liffRouter.Get("/user/:id", hdl.GetAlumniById)
	liffRouter.Post("/user", hdl.SignUp)
	liffRouter.Post("/otp/request", hdl.GetOTP)
	liffRouter.Put("/otp/verify", hdl.VerifyOTP)
}
