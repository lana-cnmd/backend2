package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lana-cnmd/backend2/pkg/repository/postgres"
	"github.com/lana-cnmd/backend2/types"
)

type IClientRepo interface {
	Create(client types.CreateClientRequest) (int, error)
	SearchClientByName(firstName, lastName string) (types.SearchClientResponse, error)
	DeleteClient(clientId int) error
	GetAllClients(limit, offset int) ([]types.SearchClientResponse, error)
	UpdateClientAddress(clientId int, newAddress types.UpdateAddressInput) error
}

type ISupplierRepo interface {
	Create(supplier types.SupplierDTO) (int, error)
	GetSupplierById(supplierId int) (types.SupplierDTO, error)
	GetAllSuppliers() ([]types.SupplierDTO, error)
	DeleteSupplierById(supplierId int) error
	UpdateSupplierAddress(supplierId int, newAddress types.UpdateAddressInput) error
}

type IImageRepo interface {
	AddImage(fileBytes []byte) (uuid.UUID, error)
	GetImageByImageUUID(imageUUID uuid.UUID) ([]byte, error)
	DeleteImageByImageUUID(imageUUID uuid.UUID) error
	UpdateImage(imageId uuid.UUID, fileBytes []byte) error
	GetImageByProductId(productId int) ([]byte, error)
}

type IProductRepo interface {
	Create(input types.CreateProductRequest) (int, error)
	GetProductById(productId int) (types.GetProductResponce, error)
	GetAllProducts() ([]types.GetProductResponce, error)
	DeleteProductById(productId int) error
	DecreaseProductAmount(productId int, decreaseAmount int) error
}

type Repository struct {
	IClientRepo
	ISupplierRepo
	IImageRepo
	IProductRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IClientRepo:   postgres.NewClientPostgresImpl(db),
		ISupplierRepo: postgres.NewSupplierPostgresImpl(db),
		IImageRepo:    postgres.NewImagePostgresImpl(db),
		IProductRepo:  postgres.NewProductPostgresImpl(db),
	}
}
