package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
)

type HouseService struct {
	repo repository.House
}

func NewHouseService(repo repository.House) *HouseService {
	return &HouseService{
		repo: repo,
	}
}

func (s *HouseService) GetAllHouses() ([]domain.House, error) {
	return s.repo.GetAllHouses()
}
func (s *HouseService) GetHouseById(houseId int) (domain.House, error) {
	return s.repo.GetHouseById(houseId)
}
