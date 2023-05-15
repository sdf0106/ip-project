package dto

type HireAgentInput struct {
	AgentId int `json:"agent_id" validate:"required"`
	HouseId int `json:"house_id" validate:"required"`
}

func (h HireAgentInput) Validate() error {
	return validate.Struct(h)
}
