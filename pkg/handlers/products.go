package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/types"
)

// @Summary Create a new product
// @Description Creates a product with the provided details and returns its ID
// @Tags products
// @Accept json
// @Produce json
// @Param product body types.CreateProductRequest true "Product details"
// @Success 200 {object} map[string]int "{'id': 123}"
// @Failure 400 {object} myError "{'message': 'Invalid request body'}"
// @Failure 500 {object} myError "{'message': 'Failed to create product'}"
// @Router /api/v1/products [post]
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

// @Summary Get product details by ID
// @Description Retrieves a product's information by its unique identifier
// @Tags products
// @Accept json
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} types.GetProductResponce "Product details"
// @Failure 400 {object} myError "{'message': 'Invalid product ID'}"
// @Failure 404 {object} myError "{'message': 'Product not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/products/{id} [get]
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

// @Description Product list response container
// @Name GetAllProductsResponse
// @Id GetAllProductsResponse
// @Property data type array items=GetProductResponce description="List of products"
type getAllProductsResponse struct {
	Data []types.GetProductResponce `json:"data"`
}

// @Summary Get all products
// @Description Returns a list of all products available in the system
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} getAllProductsResponse "List of products"
// @Failure 500 {object} myError "{'message': 'Failed to retrieve products'}"
// @Router /api/v1/products [get]
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

// @Summary Delete a product
// @Description Removes a product from the database by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} map[string]string "{'Status': 'ok'}"
// @Failure 400 {object} myError "{'message': 'Invalid product ID'}"
// @Failure 404 {object} myError "{'message': 'Product not found'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/products/{id} [delete]
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

// @Summary Decrease product stock
// @Description Reduces the available stock of a product by a specified amount
// @Tags products
// @Accept json
// @Produce json
// @Param id path integer true "Product ID"
// @Param decrease_amount body integer true "Amount to decrease (minimum 1)" example=10
// @Success 200 {object} map[string]string "{'message': 'Product amount decreased successfully'}"
// @Failure 400 {object} myError "{'message': 'Invalid request parameters'}"
// @Failure 404 {object} myError "{'message': 'Product not found'}"
// @Failure 409 {object} myError "{'message': 'Not enough stock to decrease'}"
// @Failure 500 {object} myError "{'message': 'Internal server error'}"
// @Router /api/v1/products/{id}/decrease-amount [put]
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
