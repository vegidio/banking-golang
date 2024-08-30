package auth

import (
	. "banking/internal/apigateway/shared"
	"banking/internal/apigateway/shared/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupController(router *gin.RouterGroup) {
	router.POST("/signin", postSignIn)
}

// region - Private functions

func postSignIn(c *gin.Context) {
	var dto SigninDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		Abort(c, errors.BadRequest(Params{}))
		return
	}

	tokens, err := SignIn(dto.Email, dto.Password)
	if err != nil {
		Abort(c, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// endregion
