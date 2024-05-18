package artbook

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"yanisdib/ostellers/config"
	"yanisdib/ostellers/errors"
	"yanisdib/ostellers/product"
)

var artbookCollection *mongo.Collection = config.GetCollection("artbooks")

func CreateArtbook(ctx context.Context, input *CreateInput) (*mongo.InsertOneResult, error) {

	newArtbook := toArtbook(input)

	result, err := artbookCollection.InsertOne(ctx, newArtbook)
	if err != nil {
		defer log.Println(errors.ERR_ITEM_NOT_CREATED)
		return nil, err
	}

	newArtbook.ID = result.InsertedID.(primitive.ObjectID)

	return result, nil

}

func GetArtbookByID(ctx context.Context, id string) (*Artbook, error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	var artbook *Artbook

	err := artbookCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&artbook)
	if err != nil {
		log.Fatal(errors.ERR_ITEM_NOT_FOUND)
		return nil, nil
	}

	return artbook, nil

}

func toArtbook(input *CreateInput) *Artbook {

	releaseDate, err := time.Parse("2006-01-02", input.ReleasedAt)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var formats []product.ProductFormat
	for _, format := range input.Formats {
		parsedFormat := product.ProductFormats
		formats = append(formats, parsedFormat[format])
	}

	product := &product.Product{
		Reference:    input.Reference,
		Label:        input.Label,
		Description:  input.Description,
		Categories:   input.Categories,
		Tags:         input.Tags,
		Artists:      input.Artists,
		Editors:      input.Editors,
		Pictures:     input.Pictures,
		Stock:        input.Stock,
		Price:        input.Price,
		Availability: product.Availabilities[input.Availability],
		Formats:      formats,
		ReleasedAt:   releaseDate.UTC(),
		CreatedAt:    time.Now().UTC(),
	}

	return &Artbook{
		Product:    product,
		PagesCount: input.PagesCount,
	}

}
