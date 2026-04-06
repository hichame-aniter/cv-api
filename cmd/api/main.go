package main

import (
	"cv-api/internal/config"
	"cv-api/internal/routes"
	"log"
)

func main() {

	config.LoadConfig()

	router := routes.SetupRouter()

	// Start the server on the configured port
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}
