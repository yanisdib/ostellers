package models

// TODO: add value objects
type Product struct {
	ID           string
	Label        string
	Description  string
	Categories   map[int]Category
	Quantity     uint16
	Availability Availability
	Format       ProductFormat
	ReleasedAt   string
	CreatedAt    string
	EditedAt     string
}
