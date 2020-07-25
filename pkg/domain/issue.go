package domain

import (
	"time"
)

// Issue entity
type Issue struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	ProjectID   uint      `json:"projectId"`
	Project     Project   `json:"project"`
	Labels      []Label   `json:"labels" gorm:"many2many:issues_labels;"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
