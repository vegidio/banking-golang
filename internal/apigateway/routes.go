package apigateway

import (
	"banking/internal/apigateway/users"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	userController := router.Group("/users")
	users.SetupController(userController)
}
