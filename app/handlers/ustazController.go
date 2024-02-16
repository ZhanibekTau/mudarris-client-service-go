package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	requestsForm "user-service-go/app/requests"
	"user-service-go/app/responses/errorResponse"
	"user-service-go/model"
)

func (h *Handler) createUstaz(c *gin.Context) {
	var input model.Ustaz
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.manager.CreateUstaz(&input)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ustaz was created successfully"})
}

func (h *Handler) loginUstaz(c *gin.Context) {
	var login requestsForm.Login
	if err := c.BindJSON(&login); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ustaz, err := h.manager.LoginUstaz(login, h.appData)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ustaz_id": ustaz.Id,
		"token":    ustaz.Token,
	})
}

func (h *Handler) updateUstaz(c *gin.Context) {
	var inputForUpdate model.Ustaz
	if err := c.BindJSON(&inputForUpdate); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ustaz, err := h.manager.UpdateUstaz(inputForUpdate)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ustaz_id": ustaz.Id,
	})
}

func (h *Handler) getByIdUstaz(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}

	ustaz, err := h.manager.GetByIdUstaz(idInt)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name":         ustaz.Name,
		"lastname":     ustaz.Lastname,
		"phone":        ustaz.Phone,
		"degree":       ustaz.Degree,
		"university":   ustaz.University,
		"personalInfo": ustaz.AdditionalInfo,
		"experience":   ustaz.Experience,
	})
}

func (h *Handler) requestForCreate(c *gin.Context) {
	var input model.Ustaz
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.manager.RequestForCreate(&input, h.appData)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "request was sent successfully"})
}

func (h *Handler) getByIdUstazAndUstaz(c *gin.Context) {
	var input requestsForm.ClientAndUstazId
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.manager.GetUstazAndClientById(input)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) getClientsByIds(c *gin.Context) {
	var input requestsForm.ClientIds
	if err := c.BindJSON(&input); err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.manager.GetClientsByIds(input.Ids)
	if err != nil {
		errorResponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
