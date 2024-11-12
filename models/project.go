package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	MemeberIDs []uuid.UUID `json:"member_ids" gorm:"type:json"`

	OwnerID uuid.UUID `json:"owner_id"`
	Owner User `json:"owner" gorm:"foreignKey:OwnerID"`
}