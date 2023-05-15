package dto

type HouseInCartInput struct {
	HouseId int `json:"house_id" validate:"requried"`
}

func (h HouseInCartInput) Validate() error {
	return validate.Struct(h)
}
