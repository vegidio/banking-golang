package shared

import "github.com/gin-gonic/gin"

type Params struct {
	Type     string
	Title    string
	Detail   string
	Instance string
}

func Abort(c *gin.Context, err *ApiError) {
	_ = c.Error(err)
	c.Status(err.Status)
	c.Abort()
}
