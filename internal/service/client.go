package service

import "github.com/sdf0106/ip-project/internal/repository"

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
