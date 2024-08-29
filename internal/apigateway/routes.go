package apigateway

import (
	"banking/internal/apigateway/feat/auth"
	"banking/internal/apigateway/feat/users"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	authController := router.Group("/auth")
	auth.SetupController(authController)

	userController := router.Group("/users")
	users.SetupController(userController)
}
