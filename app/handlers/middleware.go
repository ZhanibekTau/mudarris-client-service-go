package handlers

import (
	"client-service-go/app/utils"
	logger "client-service-go/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func UserIdentity(c *gin.Context) {
	log := &logger.Log{}

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Error("empty auth header", http.StatusUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		log.Error("Invalid auth header", http.StatusUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := utils.ParseToken(headerParts[1])
	if err != nil {
		log.Error("error while parsing token", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, "error while parsing token")
	}

	c.Set("userId", userId)
}
