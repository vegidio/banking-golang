package internal

import (
	"github.com/gin-gonic/gin"
	"template-golang/internal/users"
)

func SetupRoutes(router *gin.RouterGroup) {
	userController := router.Group("/users")
	users.SetupController(userController)
}
