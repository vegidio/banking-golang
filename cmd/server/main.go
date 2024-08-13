package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"template-golang/internal"
)

func main() {
	engine := gin.Default()

	// API /v1
	router := engine.Group("/v1")
	internal.SetupRoutes(router)

	err := engine.Run(":3000")
	if err != nil {
		slog.Error("Error starting server: %v", err)
	}
}
