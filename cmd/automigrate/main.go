package main

import (
	"log"

	"gitlab.com/qr-through/entry/backend/config"
	"gitlab.com/qr-through/entry/backend/infrastructure"
	"gitlab.com/qr-through/entry/backend/internal/core/model"
)

func init() {
	config.InitConfig()
	infrastructure.InitDB()
}

func main() {
	if err := infrastructure.DB.Migrator().AutoMigrate(
		model.Account{},
		model.Alumni{},
		model.Alumni_new{},
		model.Moderator{},
		model.OTPTransaction{},
		model.QRCodeTransaction{},
		model.UsageLog{},
		model.User{}); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Migrated")
}
