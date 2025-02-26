package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

func (h *Handler) addProduct(c *gin.Context) {
	var input types.CreateProductRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.IProductService.Create(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.IProductService.GetProductById(productId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

type getAllProductsResponse struct {
	Data []types.GetProductResponce `json:"data"`
}

func (h *Handler) getAllProducts(c *gin.Context) {

	products, err := h.services.IProductService.GetAllProducts()
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: products,
	})
}

func (h *Handler) deleteProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IProductService.DeleteProductById(productId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})

}

func (h *Handler) decreaseProductAmount(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil || productId <= 0 {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		DecreaseAmount int `json:"decrease_amount" binding:"required,gte=1"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IProductService.DecreaseProductAmount(productId, input.DecreaseAmount)
	if err != nil {
		if strings.Contains(err.Error(), "not enough stock") {
			newErrorResponce(c, http.StatusConflict, "Not enough stock to decrease")
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Product amount decreased successfully"})
}
