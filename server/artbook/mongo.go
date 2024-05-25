package artbook

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
		log.Println(errors.ERR_ITEM_NOT_FOUND)
		return nil, err
	}

	return artbook, nil

}

func DeleteArtbookByID(ctx context.Context, id string) (deleteCount int64, err error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	res, err := artbookCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Fatalf("An error occured while deleting artbook with ID: %s", id)
		return 0, err
	}

	if res.DeletedCount == 0 {
		log.Println(errors.ERR_ITEM_NOT_FOUND)
	}

	return res.DeletedCount, nil

}

func UpdateArtbookByID(ctx context.Context, id string, update interface{}) *Artbook {

	objectID, _ := primitive.ObjectIDFromHex(id)
	var updatedArtbook Artbook

	err := artbookCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedArtbook)

	if err != nil {
		log.Print(err)
		return nil
	}

	log.Print(&updatedArtbook.UpdatedAt)

	return &updatedArtbook

}

func toArtbook(input *CreateInput) *Artbook {

	releaseDate, err := time.Parse("2006-01-02", input.ReleasedAt)
	if err != nil {
		log.Println(err)
		return nil
	}

	releaseDate = releaseDate.UTC()
	createdAt := time.Now().UTC()

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
		ReleasedAt:   &releaseDate,
		CreatedAt:    &createdAt,
	}

	return &Artbook{
		Product:    product,
		PagesCount: input.PagesCount,
	}

}
