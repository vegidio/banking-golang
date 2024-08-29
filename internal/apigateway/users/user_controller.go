package users

import (
	. "banking/internal/apigateway/shared"
	"github.com/gin-gonic/gin"
)

func SetupController(router *gin.RouterGroup) {
	router.GET("/me", func(c *gin.Context) {
		user, _ := FindById(1)
		response := Response[UserDto]{Data: *user}
		c.JSON(200, response)
	})
}
