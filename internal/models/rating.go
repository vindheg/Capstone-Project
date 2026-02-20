package models

import "github.com/google/uuid"

type Rating struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID
	RecipeID uuid.UUID
	Score    int
}