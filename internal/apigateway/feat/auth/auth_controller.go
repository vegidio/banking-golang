package auth

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := SignIn(dto.Email, dto.Password)
	if err != nil {
		_ = c.AbortWithError(err.Status, err)
	}

	c.JSON(http.StatusOK, tokens)
}

// endregion
