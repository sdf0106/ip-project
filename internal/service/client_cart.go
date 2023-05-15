package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
)

type ClientCartService struct {
	repo repository.ClientCart
}

func NewClientCartService(repo repository.ClientCart) *ClientCartService {
	return &ClientCartService{
		repo: repo,
	}
}

func (s *ClientCartService) GetCart(clientId int) ([]domain.House, error) {
	return s.repo.GetCart(clientId)
}

func (s *ClientCartService) AddToCart(clientId int, houseId int) error {
	return s.repo.AddToCart(clientId, houseId)
}

func (s *ClientCartService) RemoveFromCart(clientId int, houseId int) error {
	return s.repo.RemoveFromCart(clientId, houseId)
}
