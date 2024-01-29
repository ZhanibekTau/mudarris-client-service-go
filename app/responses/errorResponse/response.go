package errorResponse

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.Set("exception", message)
	c.Status(statusCode)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
