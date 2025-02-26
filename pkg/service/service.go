package service

import (
	"github.com/google/uuid"
	"github.com/lana-cnmd/backend2/pkg/repository"
	"github.com/lana-cnmd/backend2/types"
)

type Service struct {
	IClientService
	ISupplierService
	IImageService
	IProductService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IClientService:   NewClientService(repos.IClientRepo),
		ISupplierService: NewSupplierService(repos.ISupplierRepo),
		IImageService:    NewImageService(repos.IImageRepo),
		IProductService:  NewProductService(repos.IProductRepo),
	}
}

type IClientService interface {
	Create(client types.CreateClientRequest) (int, error)
	SearchClientByName(firstName, lastName string) (types.SearchClientResponse, error)
	DeleteClient(clientId int) error
	GetAllClients(limit, offset int) ([]types.SearchClientResponse, error)
	UpdateClientAddress(clientId int, newAddress types.UpdateAddressInput) error
}

type ISupplierService interface {
	Create(supplier types.SupplierDTO) (int, error)
	GetSupplierById(supplierId int) (types.SupplierDTO, error)
	GetAllSuppliers() ([]types.SupplierDTO, error)
	DeleteSupplierById(supplierId int) error
	UpdateSupplierAddress(supplierId int, newAddress types.UpdateAddressInput) error
}

type IImageService interface {
	AddImage(fileBytes []byte) (uuid.UUID, error)
	GetImageByImageUUID(imageUUID uuid.UUID) ([]byte, error)
	DeleteImageByImageUUID(imageUUID uuid.UUID) error
	UpdateImage(imageId uuid.UUID, fileBytes []byte) error
	GetImageByProductId(productId int) ([]byte, error)
}

type IProductService interface {
	Create(input types.CreateProductRequest) (int, error)
	GetProductById(productId int) (types.GetProductResponce, error)
	GetAllProducts() ([]types.GetProductResponce, error)
	DeleteProductById(productId int) error
	DecreaseProductAmount(productId int, decreaseAmount int) error
}
