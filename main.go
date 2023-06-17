package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/qr-through/entry/backend/api"
	"gitlab.com/qr-through/entry/backend/config"
	"gitlab.com/qr-through/entry/backend/infrastructure"
)

var cfg config.LocalConfig

func init() {
	// config
	config.InitConfig()
	cfg = config.Config

	// lineBot
	infrastructure.InitLineBot()

	// postgresql
	infrastructure.InitDB()
}

func main() {
	app := fiber.New()
	api.InitAPI(app)

	addr := getAddress()
	log.Printf("%v started at %v", cfg.Name, cfg.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getAddress() string {
	addr := ":8000"
	if cfg.Port != "" {
		addr = fmt.Sprintf(":%v", cfg.Port)
	}
	return addr
}
