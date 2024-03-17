package category

import "github.com/google/uuid"

// TODO: add value objects
type Category struct {
	ID            uuid.UUID        `json:"id"`
	Label         string           `json:"label"`
	Description   string           `json:"description"`
	Subcategories map[int]Category `json:"subcategories"`
	CreatedAt     string           `json:"createdAt"`
}
