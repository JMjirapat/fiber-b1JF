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
	user := model.Alumni{
		ID:        620610023,
		Firstname: "Jedsadaporn",
		Lastname:  "Juntong",
		Tel:       "0981744055",
	}
	if err := infrastructure.DB.Create(&user); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Created")
}
