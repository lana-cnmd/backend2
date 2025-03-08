package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lana-cnmd/backend2/pkg/service"

	_ "github.com/lana-cnmd/backend2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		clients := api.Group("/clients")
		{
			clients.POST("/", h.addClient)
			clients.DELETE("/:id", h.deleteClient)
			clients.GET("/search", h.searchClientByName)
			clients.GET("/", h.getAllClients)
			clients.PUT("/:id/address", h.updateClientAddress)
		}

		products := api.Group("/products")
		{
			products.POST("/", h.addProduct)
			products.PUT("/:id/decrease-amount", h.decreaseProductAmount)
			products.GET("/:id", h.getProductById)
			products.GET("/", h.getAllProducts)
			products.DELETE("/:id", h.deleteProductById)
		}

		suppliers := api.Group("/suppliers")
		{
			suppliers.POST("/", h.addSupplier)
			suppliers.PUT("/:id/address", h.updateSupplierAddress)
			suppliers.DELETE("/:id", h.deleteSupplierById)
			suppliers.GET("/", h.getAllSuppliers)
			suppliers.GET("/:id", h.getSupplierById)
		}

		images := api.Group("/images")
		{
			images.POST("/", h.addImage)
			images.PUT("/:id", h.updateImage)
			images.DELETE("/:id", h.deleteImageByImageUUID)
			images.GET("/product/:id", h.getImageByProductId)
			images.GET("/:id", h.getImageByImageUUID)
		}
	}

	return router
}
