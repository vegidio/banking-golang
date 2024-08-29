package middlewares

import (
	. "banking/internal/apigateway/shared"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Creates a new ResponseWriter to capture the response body
		writer := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Process the request
		c.Next()

		// If there's an error, modify the response accordingly
		if len(c.Errors) > 0 {
			response := Response[interface{}]{
				Error: fmt.Errorf(c.Errors.String()),
			}

			c.JSON(c.Writer.Status(), response)
			return
		}

		// Attempt to unmarshal the response into a generic interface{}
		var responseData interface{}
		if err := json.Unmarshal(writer.body.Bytes(), &responseData); err != nil {
			c.JSON(http.StatusInternalServerError,
				Response[interface{}]{Error: fmt.Errorf("failed to parse response data: %v", err)})
			return
		}

		// Wrap the response in the custom Response struct
		response := Response[interface{}]{
			Data:  responseData,
			Error: nil,
		}

		// Reset the response body, otherwise it will write the original response and the new response
		c.Writer = writer.ResponseWriter

		// Send the wrapped response
		c.JSON(c.Writer.Status(), response)
	}
}

// bodyWriter is a custom ResponseWriter to capture the response body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	// Only write to the buffer, not to the response
	return w.body.Write(b)
}
