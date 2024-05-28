package product

import (
	"log"
	"time"
)

// Product defines common product attributes
type Product struct {
	Reference    string              `bson:"reference" json:"reference"`
	Label        string              `bson:"label" json:"label"`
	Description  string              `bson:"description,omitempty" json:"description,omitempty"`
	Categories   []string            `bson:"categories" json:"categories"`
	Tags         []string            `bson:"tags,omitempty" json:"tags,omitempty"`
	Artists      []string            `bson:"artists" json:"artists"`
	Editors      []string            `bson:"editors" json:"editors"`
	Pictures     []ProductImageAttrs `bson:"pictures,omitempty" json:"pictures,omitempty"`
	Stock        uint32              `bson:"stock" json:"stock"`
	Price        float32             `bson:"price,omitempty" json:"price,omitempty"`
	Availability string              `bson:"availability" json:"availability"`
	Formats      []string            `bson:"formats" json:"formats"`
	ReleasedAt   *time.Time          `bson:"released_at" json:"releasedAt"`
	CreatedAt    *time.Time          `bson:"created_at" json:"createdAt,omitempty"`
	UpdatedAt    *time.Time          `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}

// ProductImageAttrs defines a product images attributes
type ProductImageAttrs struct {
	URI     string `bson:"uri" json:"uri" validate:"required"`
	Title   string `bson:"title" json:"title" validate:"required"`
	Caption string `bson:"caption,omitempty" json:"caption,omitempty"`
	Order   uint8  `bson:"order" json:"order" validate:"required"`
}

// Availabity enumerates a product stock status
type Availability int

const (
	InStock Availability = iota
	SoldOut
	Restocking
	Limited
	UpcomingRelease
)

var AvailabilitiesByLabel = map[string]Availability{
	"in stock":         InStock,
	"sold out":         SoldOut,
	"restocking":       Restocking,
	"limited":          Limited,
	"upcoming release": UpcomingRelease,
}

// ProductFormat enumerates all different product format
type ProductFormat int

const (
	Physical ProductFormat = iota
	Digital
)

var FormatsByLabel = map[string]ProductFormat{
	"physical": Physical,
	"digital":  Digital,
}

func IsValidFormat(value string) bool {

	_, found := FormatsByLabel[value]
	if !found {
		log.Print("Invalid product format")
		return found
	}

	return found

}
