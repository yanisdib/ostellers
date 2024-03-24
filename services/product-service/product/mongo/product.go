package mongo

import (
	"context"
	"service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID           primitive.ObjectID       `bson:"_id,omitempty"`
	Label        string                   `bson:"label"`
	Description  string                   `bson:"description"`
	Stock        uint16                   `bson:"stock"`
	Availability string                   `bson:"availability"`
	Formats      []string                 `bson:"formats"`
	CreatedAt    primitive.DateTime       `bson:"createdAt"`
	UpdatedAt    primitive.DateTime       `bson:"updatedAt"`
	ReleasedAt   primitive.DateTime       `bson:"releasedAt"`
	Categories   map[int]*models.Category `bson:"categories"`
}

type ProductCollection struct {
	collection *mongo.Collection
}

func InitProductCollection(db *mongo.Database, collection string) *ProductCollection {
	return &ProductCollection{
		collection: db.Collection(collection),
	}
}

func (c ProductCollection) CreateProduct(ctx context.Context, product *models.Product) error {
	model := toModel(product)
	res, err := c.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	product.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (c ProductCollection) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	uid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "id", Value: uid}}

	var product Product
	err := c.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		panic(err)
	}

	return toProduct(&product), nil
}

func (c ProductCollection) GetProducts(ctx context.Context) ([]*models.Product, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	products := make([]*Product, 0)

	for cursor.Next(ctx) {
		product := new(Product)
		err := cursor.Decode(product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toProducts(products), nil
}

func toModel(product *models.Product) *Product {
	stringFormats := make([]string, len(product.Formats))
	for _, format := range product.Formats {
		stringFormats = append(stringFormats, string(format))
	}

	return &Product{
		Label:        product.Label,
		Description:  product.Description,
		Stock:        product.Stock,
		Availability: string(product.Availability),
		Formats:      stringFormats,
	}
}

func toProduct(product *Product) *models.Product {
	formats := make([]models.ProductFormat, len(product.Formats))
	for _, format := range product.Formats {
		pf := models.StringToProductFormat(format)
		formats = append(formats, pf)
	}

	return &models.Product{
		ID:           product.ID.Hex(),
		Label:        product.Label,
		Description:  product.Description,
		Stock:        product.Stock,
		Availability: models.Availability(product.Availability),
		Formats:      formats,
		ReleasedAt:   product.ReleasedAt.Time(),
		CreatedAt:    product.CreatedAt.Time(),
		UpdatedAt:    product.UpdatedAt.Time(),
		Categories:   product.Categories,
	}
}

func toProducts(products []*Product) []*models.Product {
	output := make([]*models.Product, len(products))
	for i, product := range products {
		output[i] = toProduct(product)
	}

	return output
}
