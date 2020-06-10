package domain

// Label entity
type Label struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ColorHexCode string `json:"colorHexCode"`
}
