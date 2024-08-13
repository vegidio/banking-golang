package users

import (
	"github.com/gin-gonic/gin"
	. "template-golang/internal/shared"
)

func SetupController(router *gin.RouterGroup) {
	router.GET("/me", func(c *gin.Context) {
		user, _ := FindById(1)
		response := Response[UserDto]{Data: *user}
		c.JSON(200, response)
	})
}
