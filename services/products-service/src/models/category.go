package models

// TODO: add value objects
type Category struct {
	ID            string
	Label         string
	Description   string
	Subcategories map[int]Category
	createdAt     string
}
