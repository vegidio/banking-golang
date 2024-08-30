package users

import (
	. "banking/internal/apigateway/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupController(router *gin.RouterGroup) {
	router.GET("/me", getMe)
}

// region - Private functions

func getMe(c *gin.Context) {
	user, err := RetrieveById(1)
	if err != nil {
		Abort(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// endregion
