package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-service-go/app/requests"
	"user-service-go/app/responses/errorResponse"
	"user-service-go/model"
)

func (h *Handler) create(c *gin.Context) {
	var input model.Client
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.manager.Create(&input)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client was created successfully"})
}

func (h *Handler) login(c *gin.Context) {
	var input requests.Login
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	client, err := h.manager.Login(input, h.appData)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"client_id": client.Id,
		"token":     client.Token,
	})
}

func (h *Handler) update(c *gin.Context) {
	var inputForUpdate model.Client
	if err := c.BindJSON(&inputForUpdate); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	client, err := h.manager.Update(inputForUpdate)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"client_id": client.Id,
	})
}

func (h *Handler) getById(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}

	client, err := h.manager.GetById(idInt)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name":     client.Name,
		"lastname": client.Lastname,
		"email":    client.Email,
		"phone":    client.Phone,
	})
}

func (h *Handler) getAllClientsIds(c *gin.Context) {
	clientids, err := h.manager.GetAllClientsIds()
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, clientids)
}
