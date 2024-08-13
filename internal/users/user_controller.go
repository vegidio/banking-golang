package users

import "github.com/gin-gonic/gin"

func SetupController(router *gin.RouterGroup) {
	router.GET("/me", func(c *gin.Context) {
		user, _ := FindById(1)
		c.JSON(200, user)
	})
}
