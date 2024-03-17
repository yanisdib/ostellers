package product

import (
	"service/category"
	"time"

	"github.com/google/uuid"
)

type Availability string

const (
	IN_STOCK         Availability = "IN_STOCK"
	OUT_OF_STOCK     Availability = "OUT_OF_STOCK"
	IN_PROVISION     Availability = "IN_PROVISION"
	NOT_YET_RELEASED Availability = "NOT_YET_RELEASED"
)

type ProductFormat string

const (
	PHYSICAL ProductFormat = "PHYSICAL"
	DIGITAL  ProductFormat = "DIGITAL"
)

// TODO: add value objects
type Product struct {
	ID           uuid.UUID                  `json:"id"`
	Label        string                     `json:"label"`
	Description  string                     `json:"description"`
	Categories   map[int]*category.Category `json:"categories"`
	Quantity     uint16                     `json:"quantity"`
	Availability Availability               `json:"availability"`
	Format       ProductFormat              `json:"format"`
	ReleasedAt   time.Time                  `json:"releaseAt"`
	CreatedAt    time.Time                  `json:"createdAt"`
	UpdatedAt    time.Time                  `json:"updatedAt"`
}
