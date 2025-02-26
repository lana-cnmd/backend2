package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

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

type getAllsuppliersResponce struct {
	Data []types.SupplierDTO `json:"data"`
}

func (h *Handler) getAllSuppliers(c *gin.Context) {
	suppliers, err := h.services.ISupplierService.GetAllSuppliers()
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllsuppliersResponce{
		Data: suppliers,
	})
}

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
