package models

import "github.com/google/uuid"

type Recipe struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Description string
	Ingredients string
	ImageURL    string
	UserID      uuid.UUID
}