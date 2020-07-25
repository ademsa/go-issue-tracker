package domain

import (
	"time"
)

// Label entity
type Label struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	ColorHexCode string    `json:"colorHexCode"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
