package dto

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type UserRoleInput struct {
	Role string `json:"role" validate:"required"`
}

func (u UserRoleInput) Validate() error {
	err := validate.Struct(u)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			if fieldErr.Field() == "role" && (fieldErr.Param() != "client" || fieldErr.Param() != "owner" || fieldErr.Param() != "agent") {
				return errors.New("invalid type of user")
			}
		}
	}
	return nil
}
