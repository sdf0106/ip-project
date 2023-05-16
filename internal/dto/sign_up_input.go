package dto

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"time"
)

type SignUpInput struct {
	Name     string `json:"name" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=5"`
}

func (s SignUpInput) Validate() error {
	return validate.Struct(s)
}

func (s SignUpInput) InputToEntity() domain.User {
	return domain.User{
		Email:        s.Email,
		Name:         s.Name,
		Password:     s.Password,
		RegisteredAt: time.Now(),
	}
}
