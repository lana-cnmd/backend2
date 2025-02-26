package service

import (
	"github.com/lana-cnmd/backend2/pkg/repository"
	"github.com/lana-cnmd/backend2/types"
)

type ClientService struct {
	//репозиторий с которым будем общаться
	repo repository.IClientRepo
}

func NewClientService(repo repository.IClientRepo) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(client types.CreateClientRequest) (int, error) {
	return s.repo.Create(client)
}

func (s *ClientService) SearchClientByName(firstName, lastName string) (types.SearchClientResponse, error) {
	return s.repo.SearchClientByName(firstName, lastName)
}

func (s *ClientService) DeleteClient(clientId int) error {
	return s.repo.DeleteClient(clientId)
}

func (s *ClientService) GetAllClients(limit, offset int) ([]types.SearchClientResponse, error) {
	return s.repo.GetAllClients(limit, offset)
}

func (s *ClientService) UpdateClientAddress(clientId int, newAddress types.UpdateAddressInput) error {
	return s.repo.UpdateClientAddress(clientId, newAddress)
}
