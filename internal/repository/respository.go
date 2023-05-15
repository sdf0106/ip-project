package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type Repositories struct {
	Authentication
	Agent
	Client
	ClientCart
	House
	Owner
	User
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Authentication: NewAuthenticationPostgres(db),
		ClientCart:     NewClientCartPostgres(db),
		House:          NewHousePostgres(db),
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
		GetUser(email, password string) (domain.User, error)
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
