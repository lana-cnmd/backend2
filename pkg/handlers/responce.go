package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type myError struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	fmt.Printf("send error responce with message: %s/n", message)
	c.AbortWithStatusJSON(statusCode, myError{Message: message})
}
