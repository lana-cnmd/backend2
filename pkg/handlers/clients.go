package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

// @Summary Create a new client
// @Description Creates a client with the provided personal details and returns its ID
// @Tags clients
// @Accept json
// @Produce json
// @Param client body types.CreateClientRequest true "Client details"
// @Success 200 {object} map[string]int "{'id': 123}"
// @Failure 400 {object} myError "{'message': 'Invalid request body'}"
// @Failure 500 {object} myError "{'message': 'Failed to create client'}"
// @Router /api/v1/clients [post]
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

// @Summary Search client by name
// @Description Finds a client by their first and last name
// @Tags clients
// @Accept json
// @Produce json
// @Param first_name query string true "Client's first name"
// @Param last_name query string true "Client's last name"
// @Success 200 {object} types.SearchClientResponse "Client details"
// @Failure 400 {object} myError "{'message': 'empty first or last name'}"
// @Failure 404 {object} myError "{'message': 'Client not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/clients/search [get]
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

// @Summary Delete a client
// @Description Removes a client from the database by their ID
// @Tags clients
// @Accept json
// @Produce json
// @Param id path integer true "Client ID"
// @Success 200 {object} map[string]string "{'Status': 'ok'}"
// @Failure 400 {object} myError "{'message': 'Invalid client ID'}"
// @Failure 404 {object} myError "{'message': 'Client not found'}"
// @Failure 500 {object} myError "{'message': 'Failed to delete client'}"
// @Router /api/v1/clients/{id} [delete]
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

// @Summary Get all clients
// @Description Returns a list of all clients with optional pagination (limit/offset)
// @Tags clients
// @Accept json
// @Produce json
// @Param limit query integer false "Number of clients per page" example=10
// @Param offset query integer false "Skip first N clients" example=0
// @Success 200 {object} getAllClientsResponse "List of clients"
// @Failure 400 {object} myError "{'message': 'Invalid limit/offset values'}"
// @Failure 500 {object} myError "{'message': 'Failed to retrieve clients'}"
// @Router /api/v1/clients [get]
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

// @Description Clients list response container
// @Name GetAllClientsResponse
// @Id GetAllClientsResponse
// @Property data type array items=SearchClientResponse description="List of clients"
type getAllClientsResponse struct {
	Data []types.SearchClientResponse `json:"data"`
}

// @Summary Update client's address
// @Description Partially updates the address details of a client by their ID (at least one field must be provided)
// @Tags clients
// @Accept json
// @Produce json
// @Param id path integer true "Client ID"
// @Param address body types.UpdateAddressInput true "New address details (at least one field required)"
// @Success 200 {object} map[string]string "{'Status': 'ok'}"
// @Failure 400 {object} myError "{'message': 'Invalid request parameters'}"
// @Failure 404 {object} myError "{'message': 'Client not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/clients/{id}/address [put]
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
