package config

import (
	"time"

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

type TaskRequestBody struct {
	Title        string    `json:"title"`
	Description string    `json:"description"`
	Status string `json:"status,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	AssignedToID uuid.UUID `json:"assigned_to_id"`
}