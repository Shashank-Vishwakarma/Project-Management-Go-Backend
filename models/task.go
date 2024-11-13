package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CreatedByID uuid.UUID `json:"created_by"`
	CreatedByUser User `json:"created_by_user" gorm:"foreignKey:CreatedByID"`

	AssignedToID uuid.UUID `json:"assigned_to_id"`
	AssignedTo User `json:"assigned_to" gorm:"foreignKey:AssignedToID"`

	ProjectID uuid.UUID `json:"project_id"`
	Project Project `json:"project" gorm:"foreignKey:ProjectID"`
}