package dto

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"time"
)

type CreateHouseInput struct {
	Address   string    `json:"address" validate:"required,gte=4"`
	OwnerId   int       `json:"owner_id"`
	AgentId   int       `json:"agent_id"`
	Price     float64   `json:"price" validate:"required"`
	BuildDate time.Time `json:"build_date"`
}

func (c CreateHouseInput) Validate() error {
	return validate.Struct(c)
}

func (c CreateHouseInput) InputToEntity() domain.House {
	return domain.House{
		Address:   c.Address,
		OwnerId:   c.OwnerId,
		AgentId:   c.AgentId,
		Price:     c.Price,
		BuildDate: c.BuildDate,
	}
}
