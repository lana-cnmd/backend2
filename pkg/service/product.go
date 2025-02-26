package service

import (
	"github.com/lana-cnmd/backend2/pkg/repository"
	"github.com/lana-cnmd/backend2/types"
)

type ProductService struct {
	repo repository.IProductRepo
}

func NewProductService(repo repository.IProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(input types.CreateProductRequest) (int, error) {
	return s.repo.Create(input)
}

func (s *ProductService) GetProductById(productId int) (types.GetProductResponce, error) {
	return s.repo.GetProductById(productId)
}

func (s *ProductService) GetAllProducts() ([]types.GetProductResponce, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) DeleteProductById(productId int) error {
	return s.repo.DeleteProductById(productId)
}

func (s *ProductService) DecreaseProductAmount(productId int, decreaseAmount int) error {
	return s.repo.DecreaseProductAmount(productId, decreaseAmount)
}
