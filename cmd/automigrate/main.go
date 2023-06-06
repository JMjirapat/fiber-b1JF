package main

import (
	"fiber/config"
	"fiber/infrastructure"
	"fiber/internal/core/model"
	"log"
)

func init() {
	config.InitConfig()
	infrastructure.InitDB()
}

func main() {
	if err := infrastructure.DB.Migrator().AutoMigrate(
		model.Account{},
		model.Alumni{},
		model.Moderator{},
		model.OTPTranstaction{},
		model.QRCodeTransaction{},
		model.UsageLog{},
		model.User{}); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Migrated")
}
