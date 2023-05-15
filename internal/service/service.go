package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
	"github.com/sdf0106/ip-project/pkg/auth"
)

type Services struct {
	Agent
	Authentication
	Client
	ClientCart
	House
	Owner
	User
	auth.TokenManager
}

func NewService(repos *repository.Repositories, tokenManager auth.TokenManager) *Services {
	return &Services{
		Authentication: NewAuthenticationService(repos.Authentication, tokenManager),
		Agent:          NewAgentService(repos.Agent),
		Client:         NewClientService(repos.Client),
		ClientCart:     NewClientCartService(repos.ClientCart),
		House:          NewHouseService(repos.House),
		Owner:          NewOwnerService(repos.Owner),
		User:           NewUserService(repos.User),
		TokenManager:   tokenManager,
	}
}

type (
	Agent interface {
		GetAllAgents() ([]domain.Agent, error)
		GetAgentById(id int) (domain.Agent, error)
		GetAgentHouses(id int) ([]domain.House, error)
	}
	Authentication interface {
		CreateUser(user domain.User) (int, error)
		GenerateToken(email, password string) (string, error)
		ParseToken(accessToken string) (int, error)
	}
	Client     interface{}
	ClientCart interface {
		GetCart(clientId int) ([]domain.House, error)
		AddToCart(clientId int, houseId int) error
		RemoveFromCart(clientId int, houseId int) error
	}
	House interface {
		GetAllHouses() ([]domain.House, error)
		GetHouseById(houseId int) (domain.House, error)
	}
	Owner interface {
		GetMyHouses(userId int) ([]domain.House, error)
		CreateHouse(userId int, house domain.House) (int, error)
		DeleteHouse(ownerId int, houseId int) error
		UpdateHouse(ownerId int, house domain.House) (domain.House, error)
		HireAgent(ownerId int, houseId int, agentId int) (int, error)
	}
	User interface {
		ChooseRole(userId int, role string) error
		UpdateUserInfo(userId int, user domain.User) (domain.User, error)
		ChangeRole(userId int, role string) error
	}
)
