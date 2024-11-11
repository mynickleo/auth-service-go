package main

import (
	"log"
	"auth-service-backend/internal/app"
)

func main() {
	err := app.InitializationApp()
	if err != nil {
		log.Fatal(err)
	}
}
