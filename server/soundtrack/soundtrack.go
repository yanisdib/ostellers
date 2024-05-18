package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"yanisdib/ostellers/product"
)

// Soundtrack defines a soundtrack product
type Soundtrack struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	*product.Product
	Genres    []string `bson:"genres" json:"genres" validate:"required"`
	Tracklist []Disc   `bson:"tracklist" json:"tracklist,omitempty"`
}

// Disc defines the attributes of a soundtrack disc
type Disc struct {
	Title  string  `bson:"title" json:"title,omitempty"`
	Tracks []Track `bson:"tracks" json:"tracks" validate:"required"`
	Order  uint8   `bson:"order" json:"order" validate:"required"`
}

// Track defines the attributes of a disc track
type Track struct {
	Title string `bson:"title" json:"title" validate:"required"`
	Order uint8  `bson:"order" json:"order" validate:"required"`
}
