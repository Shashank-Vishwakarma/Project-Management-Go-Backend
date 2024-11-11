package config

import (
	"github.com/google/uuid"
)

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ProjectRequestBody struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID uuid.UUID `json:"owner_id"`
}