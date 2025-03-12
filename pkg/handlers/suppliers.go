package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

// @Summary Create a new supplier
// @Description Adds a new supplier to the system with the provided details
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body types.SupplierDTO true "Supplier details"
// @Success 200 {object} map[string]int "{'id': 123}"
// @Failure 400 {object} myError "{'message': 'Invalid request body'}"
// @Failure 500 {object} myError "{'message': 'Failed to create supplier'}"
// @Router /api/v1/suppliers [post]
func (h *Handler) addSupplier(c *gin.Context) {
	var input types.SupplierDTO
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ISupplierService.Create(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get supplier details by ID
// @Description Retrieves a supplier's information by its unique identifier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path integer true "Supplier ID"
// @Success 200 {object} types.SupplierDTO "Supplier details"
// @Failure 400 {object} myError "{'message': 'Invalid supplier ID'}"
// @Failure 404 {object} myError "{'message': 'Supplier not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/suppliers/{id} [get]
func (h *Handler) getSupplierById(c *gin.Context) {
	supplierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	supplier, err := h.services.ISupplierService.GetSupplierById(supplierId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// swagger:response getAllSuppliersResponse
// @Description Suppliers list response container
// @Name GetAllSuppliersResponse
// @Id GetAllSuppliersResponse
// @Property data type array items=SupplierDTO description="List of suppliers"
type getAllSuppliersResponce struct {
	Data []types.SupplierDTO `json:"data"`
}

func (h *Handler) getAllSuppliers(c *gin.Context) {
	suppliers, err := h.services.ISupplierService.GetAllSuppliers()
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSuppliersResponce{
		Data: suppliers,
	})
}

// @Summary Delete a supplier
// @Description Removes a supplier from the database by its ID
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path integer true "Supplier ID"
// @Success 200 {object} map[string]string "{'Status': 'ok'}"
// @Failure 400 {object} myError "{'message': 'Invalid supplier ID'}"
// @Failure 404 {object} myError "{'message': 'Supplier not found'}"
// @Failure 500 {object} myError "{'message': 'Failed to delete supplier'}"
// @Router /api/v1/suppliers/{id} [delete]
func (h *Handler) deleteSupplierById(c *gin.Context) {
	supplierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ISupplierService.DeleteSupplierById(supplierId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

// @Summary Update supplier's address
// @Description Partially updates the address details of a supplier by their ID (at least one field required)
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path integer true "Supplier ID"
// @Param address body types.UpdateAddressInput true "New address details"
// @Success 200 {object} map[string]string "{'Status': 'ok'}"
// @Failure 400 {object} myError "{'message': 'Invalid request parameters'}"
// @Failure 404 {object} myError "{'message': 'Supplier not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/suppliers/{id}/address [put]
func (h *Handler) updateSupplierAddress(c *gin.Context) {
	supplierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	var input types.UpdateAddressInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ISupplierService.UpdateSupplierAddress(supplierId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
