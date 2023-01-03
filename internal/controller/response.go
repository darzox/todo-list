package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorHandler struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorHandler{message})
}
