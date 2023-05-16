package domain

import (
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	UserType     string    `json:"user_type"`
	RegisteredAt time.Time `json:"registered_at"`
}
