package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
)

type OwnerService struct {
	repo repository.Owner
}

func NewOwnerService(repo repository.Owner) *OwnerService {
	return &OwnerService{
		repo: repo,
	}
}

func (s *OwnerService) GetMyHouses(userId int) ([]domain.House, error) {
	return s.repo.GetMyHouses(userId)
}
func (s *OwnerService) CreateHouse(userId int, house domain.House) (int, error) {
	return s.repo.CreateHouse(userId, house)
}
func (s *OwnerService) DeleteHouse(ownerId int, houseId int) error {
	return s.repo.DeleteHouse(ownerId, houseId)
}
func (s *OwnerService) UpdateHouse(ownerId int, house domain.House) (domain.House, error) {
	return s.repo.UpdateHouse(ownerId, house)
}
func (s *OwnerService) HireAgent(ownerId int, houseId int, agentId int) (int, error) {
	return s.repo.HireAgent(ownerId, houseId, agentId)
}
