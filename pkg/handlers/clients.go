package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

func (h *Handler) addClient(c *gin.Context) {
	var input types.CreateClientRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.IClientService.Create(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) searchClientByName(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	if firstName == "" || lastName == "" {
		newErrorResponce(c, http.StatusBadRequest, "empty first or last name")
		return
	}

	client, err := h.services.IClientService.SearchClientByName(firstName, lastName)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) deleteClient(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IClientService.DeleteClient(clientId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})

}

func (h *Handler) getAllClients(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "")
	offsetStr := c.DefaultQuery("offset", "")

	var limit, offset int
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			newErrorResponce(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			newErrorResponce(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	// Если параметры не указаны, устанавливаем значения по умолчанию
	if limit == 0 {
		limit = 100
	}

	clients, err := h.services.IClientService.GetAllClients(limit, offset)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllClientsResponse{
		Data: clients,
	})
}

type getAllClientsResponse struct {
	Data []types.SearchClientResponse `json:"data"`
}

func (h *Handler) updateClientAddress(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	var input types.UpdateAddressInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IClientService.UpdateClientAddress(clientId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
