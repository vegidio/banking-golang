package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"template-golang/internal"
)

func main() {
	engine := gin.Default()

	// Connect to the database
	dbClient, err := internal.SetupRepositories()
	if err != nil {
		log.Fatal("Failed to connect to Postgres: ", err)
	}

	defer dbClient.Close()

	// Endpoints
	router := engine.Group("/v1")
	internal.SetupRoutes(router)

	err = engine.Run(":3000")
	if err != nil {
		slog.Error("Error starting server: %w", err)
	}
}
