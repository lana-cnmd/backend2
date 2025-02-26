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

type CreateProductRequest struct {
	Name           string          `json:"name"`
	Category       string          `json:"category"`
	Price          decimal.Decimal `json:"price"`
	AvailableStock int             `json:"available_stock"` // число закупленных экземпляров товара
	SupplierId     int             `json:"supplier_id"`
	ImageId        uuid.UUID       `json:"image_id"`
}

type GetProductResponce struct {
	Name           string          `json:"name" db:"name"`
	Category       string          `json:"category" db:"category"`
	Price          decimal.Decimal `json:"price" db:"price"`
	AvailableStock int             `json:"available_stock" db:"available_stock"`   // число закупленных экземпляров товара
	LastUpdateDate time.Time       `json:"last_update_date" db:"last_update_date"` // число последней закупки
	SupplierId     int             `json:"supplier_id" db:"supplier_id"`
	ImageId        uuid.UUID       `json:"image_id" db:"image_id"`
}
