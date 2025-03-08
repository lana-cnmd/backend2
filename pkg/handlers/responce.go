package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// myError represents an API error response
// @Description Error response object
// @Name myError
// @Id myError
// @Property message type string description="Error message" example="Product not found"
type myError struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	fmt.Printf("send error responce with message: %s/n", message)
	c.AbortWithStatusJSON(statusCode, myError{Message: message})
}
