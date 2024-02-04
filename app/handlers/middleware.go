package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"user-service-go/app/utils"
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

func (h *Handler) ustazIdentity(c *gin.Context) {
	headerParts := h.getAuthorizationHeader(c)

	userId, err := utils.ParseUstazToken(headerParts[1], h.appData)
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
