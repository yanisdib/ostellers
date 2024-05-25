package artbook

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"yanisdib/ostellers/product"
)

// Artbook defines an artbook product
type Artbook struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	*product.Product `bson:",inline" json:",inline"`
	PagesCount       uint16 `bson:"pagesCount" json:"pagesCount,omitempty"`
}
