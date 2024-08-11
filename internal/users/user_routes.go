package users

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) {
	r.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Vinicius",
		})
	})
}
