package product

import (
	"time"
)

// Product defines common product attributes
type Product struct {
	Reference    string              `bson:"reference" json:"reference"`
	Label        string              `bson:"label" json:"label"`
	Description  string              `bson:"description,omitempty" json:"description,omitempty"`
	Categories   []string            `bson:"categories" json:"categories"`
	Tags         []string            `bson:"tags,omitempty" json:"tags"`
	Artists      []string            `bson:"artists" json:"artists"`
	Editors      []string            `bson:"editors" json:"editors"`
	Pictures     []ProductImageAttrs `bson:"pictures,omitempty" json:"pictures,omitempty"`
	Stock        uint32              `bson:"stock" json:"stock"`
	Price        float32             `bson:"price,omitempty" json:"price,omitempty"`
	Availability Availability        `bson:"availability" json:"availability,omitempty"`
	Formats      []ProductFormat     `bson:"formats" json:"formats,omitempty"`
	ReleasedAt   *time.Time          `bson:"released_at" json:"releasedAt,omitempty"`
	CreatedAt    *time.Time          `bson:"created_at" json:"createdAt,omitempty"`
	UpdatedAt    *time.Time          `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}

// ProductImageAttrs defines a product images attributes
type ProductImageAttrs struct {
	URI     string `bson:"uri" json:"uri" validate:"required"`
	Title   string `bson:"title" json:"title" validate:"required"`
	Caption string `bson:"caption" json:"caption,omitempty"`
	Order   uint8  `bson:"order" json:"order" validate:"required"`
}

// Availabity enumerates a product stock status
type Availability int

const (
	IN_STOCK Availability = iota
	SOLD_OUT
	RESTOCKING
	LIMITED
	UPCOMING_RELEASE
)

var Availabilities = map[string]Availability{
	"IN STOCK":         IN_STOCK,
	"SOLD OUT":         SOLD_OUT,
	"RESTOCKING":       RESTOCKING,
	"LIMITED":          LIMITED,
	"UPCOMING RELEASE": UPCOMING_RELEASE,
}

// ProductFormat enumerates all different product format
type ProductFormat int

const (
	PHYSICAL ProductFormat = iota
	DIGITAL
)

var ProductFormats = map[string]ProductFormat{
	"PHYSICAL": PHYSICAL,
	"DIGITAL":  DIGITAL,
}
