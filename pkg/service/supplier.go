package service

import (
	"github.com/lana-cnmd/backend2/pkg/repository"
	"github.com/lana-cnmd/backend2/types"
)

type SupplierService struct {
	repo repository.ISupplierRepo
}

func NewSupplierService(repo repository.ISupplierRepo) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s *SupplierService) Create(supplier types.SupplierDTO) (int, error) {
	return s.repo.Create(supplier)
}

func (s *SupplierService) GetSupplierById(supplierId int) (types.SupplierDTO, error) {
	return s.repo.GetSupplierById(supplierId)
}

func (s *SupplierService) GetAllSuppliers() ([]types.SupplierDTO, error) {
	return s.repo.GetAllSuppliers()
}

func (s *SupplierService) DeleteSupplierById(supplierId int) error {
	return s.repo.DeleteSupplierById(supplierId)
}

func (s *SupplierService) UpdateSupplierAddress(supplierId int, newAddress types.UpdateAddressInput) error {
	return s.repo.UpdateSupplierAddress(supplierId, newAddress)
}
