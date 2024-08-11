package internal

import (
	"github.com/gin-gonic/gin"
	"template-golang/internal/users"
)

func SetupRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	users.SetupRoutes(userRoutes)
}
