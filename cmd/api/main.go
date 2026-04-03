package main

import (
	"log"

	"ecommerce/internal/server"
	"ecommerce/internal/config"
	"ecommerce/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.New(cfg.DBUrl)
	
	app := server.New(cfg, db)

	err := app.Start()

	if err != nil {
		log.Fatal(err)
	}
}