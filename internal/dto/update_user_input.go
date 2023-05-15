package dto

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"time"
)

type UpdateUserInput struct {
	Name     string    `json:"name" validate:"gte=4"`
	Email    string    `json:"email" validate:"email"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday"`
	UserType string    `json:"user_type"`
}

func (s UpdateUserInput) Validate() error {
	return validate.Struct(s)
}

func (s UpdateUserInput) InputToEntity() domain.User {
	return domain.User{
		Email:        s.Email,
		Name:         s.Name,
		Birthday:     s.Birthday,
		UserType:     s.UserType,
		RegisteredAt: time.Now(),
	}
}
