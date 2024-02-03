package handlers

import (
	"client-service-go/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) clientIdentity(c *gin.Context) {
	headerParts := h.getAuthorizationHeader(c)

	userId, err := utils.ParseClientToken(headerParts[1], h.appData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "error while parsing token")
	}

	c.Set("userId", userId)
}

func (h *Handler) getAuthorizationHeader(c *gin.Context) []string {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return []string{""}
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid auth header")
		return []string{""}
	}

	return headerParts
}
