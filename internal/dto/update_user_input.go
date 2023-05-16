package dto

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"time"
)

type UpdateUserInput struct {
	Name     string `json:"name" validate:"gte=4"`
	Email    string `json:"email" validate:"email"`
	UserType string `json:"user_type"`
}

func (s UpdateUserInput) Validate() error {
	return validate.Struct(s)
}

func (s UpdateUserInput) InputToEntity() domain.User {
	return domain.User{
		Email:        s.Email,
		Name:         s.Name,
		UserType:     s.UserType,
		RegisteredAt: time.Now(),
	}
}
