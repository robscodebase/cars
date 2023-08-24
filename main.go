package main

import (
	"cars/config"
	"cars/controller"
	"cars/models"
	"cars/store"
	"log"
)

func main() {
	log.Println("main starting")
	err := config.Init(".env")
	if err != nil {
		log.Fatalf("main failed to init config: %v", err)
	}
	s, err := store.NewStore()
	if err != nil {
		log.Fatalf("main failed to create store: %v", err)
	}
	if config.Get().Postgres.RunMigrations == "true" {
		err = s.Postgres.DB.AutoMigrate(
			&models.Car{},
		).Error
		if err != nil {
			log.Fatalf("main failed to run migrations: %v", err)
		}
	}
	r := controller.InitServer(s)
	log.Println("main starting server")
	err = r.Run(config.Get().ListenPort)
	if err != nil {
		log.Fatalf("main failed to run server: %v", err)
	}
}
