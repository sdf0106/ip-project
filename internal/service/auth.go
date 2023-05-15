package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
	"github.com/sdf0106/ip-project/pkg/auth"
	"time"
)

const (
	salt     = "asdbfiwqeubflkasjdbfp9ub"
	tokenTTL = 24 * time.Hour
)

type AuthenticationService struct {
	repo repository.Authentication
	auth.TokenManager
}

func NewAuthenticationService(repo repository.Authentication, tokenManager auth.TokenManager) *AuthenticationService {
	return &AuthenticationService{
		repo:         repo,
		TokenManager: tokenManager,
	}
}

func (s *AuthenticationService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthenticationService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))

	if err != nil {
		fmt.Printf("token gereration: %s", err.Error())
		return "", err
	}

	return s.TokenManager.CreateJWT(user.Id, time.Duration(tokenTTL))
}

func (s *AuthenticationService) ParseToken(accessToken string) (int, error) {
	id, err := s.TokenManager.Parse(accessToken)

	if err != nil {
		fmt.Printf("token is not parsed: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
