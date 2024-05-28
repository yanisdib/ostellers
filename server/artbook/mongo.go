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

var artbooksCollection *mongo.Collection = config.GetCollection("artbooks")

// Create Artbook stores a new artbook in database
func CreateArtbook(ctx context.Context, input *CreateInput) (*mongo.InsertOneResult, error) {

	newArtbook := toArtbook(input)

	result, err := artbooksCollection.InsertOne(ctx, newArtbook)
	if err != nil {
		defer log.Println(errors.ERR_ITEM_NOT_CREATED)
		return nil, err
	}

	newArtbook.ID = result.InsertedID.(primitive.ObjectID)

	return result, nil

}

// GetAllArtbooks finds all artbooks stored in database
func GetAllArtbooks(ctx context.Context) []*Artbook {

	cursor, err := artbooksCollection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		log.Print(err)
		return nil
	}

	var artbooks []*Artbook
	if err := cursor.All(ctx, &artbooks); err != nil {
		log.Print(err)
		return nil
	}

	return artbooks

}

// GetArtbookByID finds an artbook for a given ID in database
func GetArtbookByID(ctx context.Context, id string) (*Artbook, error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	var artbook *Artbook

	err := artbooksCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&artbook)
	if err != nil {
		log.Println(errors.ERR_ITEM_NOT_FOUND)
		return nil, err
	}

	return artbook, nil

}

// DeleteArtbookByID deletes an artbook for a given ID in database
func DeleteArtbookByID(ctx context.Context, id string) (deleteCount int64, err error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	result, err := artbooksCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Fatalf("An error occured while deleting artbook with ID: %s", id)
		return 0, err
	}

	if result.DeletedCount == 0 {
		log.Println(errors.ERR_ITEM_NOT_FOUND)
	}

	return result.DeletedCount, nil

}

// UpdateArtbookByID finds an artbook for a given ID and update it
func UpdateArtbookByID(ctx context.Context, id string, update interface{}) *Artbook {

	objectID, _ := primitive.ObjectIDFromHex(id)

	var updatedArtbook Artbook

	err := artbooksCollection.FindOneAndUpdate(
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

// toArtbook converts a create input to an artbook struct
func toArtbook(input *CreateInput) *Artbook {

	releaseDate, err := time.Parse("2006-01-02", input.ReleasedAt)
	if err != nil {
		log.Println(err)
		return nil
	}

	releaseDate = releaseDate.UTC()
	createdAt := time.Now().UTC()

	for _, format := range input.Formats {
		isValid := product.IsValidFormat(format)
		if !isValid {
			return nil
		}
	}

	_, found := product.AvailabilitiesByLabel[input.Availability]
	if !found {
		log.Print("Invalid availability")
		return nil
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
		Availability: input.Availability,
		Formats:      input.Formats,
		ReleasedAt:   &releaseDate,
		CreatedAt:    &createdAt,
	}

	return &Artbook{
		Product:    product,
		PagesCount: input.PagesCount,
	}

}
