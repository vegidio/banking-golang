package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupController(router *gin.RouterGroup) {
	router.GET("/me", getMe)
}

// region - Private functions

func getMe(c *gin.Context) {
	user, _ := RetrieveById(1)
	c.JSON(http.StatusOK, user)
}

// endregion
