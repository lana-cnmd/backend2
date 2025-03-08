package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	Category       string          `json:"category"`
	Price          decimal.Decimal `json:"price"`
	AvailableStock int             `json:"available_stock"`  // число закупленных экземпляров товара
	LastUpdateDate time.Time       `json:"last_update_date"` // число последней закупки
	SupplierId     int             `json:"supplier_id"`
	ImageId        uuid.UUID       `json:"image_id"`
}

// CreateProductRequest represents the request to create a new product
// @Description Product creation request object
// @Name CreateProductRequest
// @Id CreateProductRequest
// @Property name type string description="Product name" example="Wireless Headphones" required=true
// @Property category type string description="Product category" example="Electronics" required=true
// @Property price type number description="Product price (e.g., 199.99)" example="199.99" required=true format=number
// @Property availableStock type integer description="Initial stock quantity" example=100 required=true
// @Property supplierId type integer description="Supplier ID" example=1 required=true
// @Property imageId type string description="Image UUID (optional)" format=uuid
type CreateProductRequest struct {
	Name           string          `json:"name"`
	Category       string          `json:"category"`
	Price          decimal.Decimal `json:"price"`
	AvailableStock int             `json:"available_stock"` // число закупленных экземпляров товара
	SupplierId     int             `json:"supplier_id"`
	ImageId        uuid.UUID       `json:"image_id"`
}

// GetProductResponce represents product details
// @Description Product details response object
// @Name GetProductResponce
// @Id GetProductResponce
// @Property id type integer description="Unique product ID" example=123
// @Property name type string description="Product name" example="Wireless Headphones"
// @Property category type string description="Product category" example="Electronics"
// @Property price type number description="Product price" example="199.99"
// @Property availableStock type integer description="Current stock quantity" example=95
// @Property lastUpdateDate type string description="Last purchase date" example="2023-10-01T12:00:00Z" format=date-time
// @Property supplierId type integer description="Supplier ID" example=1
// @Property imageId type string description="Image UUID" format=uuid
type GetProductResponce struct {
	Name           string          `json:"name" db:"name"`
	Category       string          `json:"category" db:"category"`
	Price          decimal.Decimal `json:"price" db:"price"`
	AvailableStock int             `json:"available_stock" db:"available_stock"`   // число закупленных экземпляров товара
	LastUpdateDate time.Time       `json:"last_update_date" db:"last_update_date"` // число последней закупки
	SupplierId     int             `json:"supplier_id" db:"supplier_id"`
	ImageId        uuid.UUID       `json:"image_id" db:"image_id"`
}
