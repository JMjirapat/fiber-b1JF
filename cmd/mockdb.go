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
	user := model.Alumni{
		ID:        620610023,
		Firstname: "Jedsadaporn",
		Lastname:  "Juntong",
	}
	if err := infrastructure.DB.Create(&user); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Created")
}
