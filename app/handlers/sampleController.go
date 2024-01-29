package handlers

import (
	"client-service-go/app/responses/errorResponse"
	"client-service-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//TODO:Check user by email if doesnt exists then create

	c.JSON(http.StatusOK, gin.H{"message": input})
}

func (h *Handler) login(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//TODO:Check user by email if doesnt exists then create

	c.JSON(http.StatusOK, gin.H{"message": input})
}
