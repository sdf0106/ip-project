package dto

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"time"
)

type UpdateHouseInput struct {
	Address   string    `json:"address"`
	AgentId   int       `json:"agent_id"`
	OwnerId   int       `json:"owner_id"`
	Price     float64   `json:"price"`
	BuildDate time.Time `json:"build_date"`
}

func (u UpdateHouseInput) Validate() error {
	return validate.Struct(u)
}

func (u UpdateHouseInput) InputToEntity() domain.House {
	return domain.House{
		Address:   u.Address,
		OwnerId:   u.OwnerId,
		AgentId:   u.AgentId,
		Price:     u.Price,
		BuildDate: u.BuildDate,
	}
}
