package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	OwnerID uuid.UUID `json:"owner_id"`
	Owner User `json:"owner" gorm:"foreignKey:OwnerID"`
}