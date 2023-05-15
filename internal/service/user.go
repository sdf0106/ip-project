package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) ChooseRole(userId int, role string) error {
	return s.repo.ChooseRole(userId, role)
}
func (s *UserService) UpdateUserInfo(userId int, user domain.User) (domain.User, error) {
	return s.repo.UpdateUserInfo(userId, user)
}
func (s *UserService) ChangeRole(userId int, prevRole string, role string) error {
	return s.repo.ChangeRole(userId, prevRole, role)
}
