package main

import (
	. "banking/internal/apigateway"
	"banking/internal/apigateway/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

func main() {
	engine := gin.Default()

	// Initialize middlewares
	engine.Use(middlewares.ResponseMiddleware())

	// Connect to the database
	dbClient, err := SetupRepositories()
	if err != nil {
		log.Fatal("Failed to connect to Postgres: ", err)
	}

	defer dbClient.Close()

	// Endpoints
	router := engine.Group("/api/v1")
	SetupRoutes(router)

	err = engine.Run(":3001")
	if err != nil {
		slog.Error("Error starting server: %w", err)
	}
}
