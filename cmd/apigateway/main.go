package main

import (
	"banking/internal/apigateway"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

func main() {
	engine := gin.Default()

	// Connect to the database
	dbClient, err := apigateway.SetupRepositories()
	if err != nil {
		log.Fatal("Failed to connect to Postgres: ", err)
	}

	defer dbClient.Close()

	// Endpoints
	router := engine.Group("/v1")
	apigateway.SetupRoutes(router)

	err = engine.Run(":3001")
	if err != nil {
		slog.Error("Error starting server: %w", err)
	}
}
