package main

import (
	"log"

	"github.com/Okenamay/urlshortener/internal/app/configs"
	"github.com/Okenamay/urlshortener/internal/app/database"
	"github.com/Okenamay/urlshortener/internal/app/server"
)

func main() {
	database.InitDB()

	log.Printf("Starting server on port %s", configs.ServerPort)
	server.Launch()
}
