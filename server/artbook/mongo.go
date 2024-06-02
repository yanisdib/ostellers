package artbook

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"yanisdib/ostellers/config"
	"yanisdib/ostellers/product"
)

var artbooksCollection *mongo.Collection = config.GetCollection("artbooks")

// Create Artbook stores a new artbook in database
func CreateArtbook(ctx context.Context, input *CreateInput) (*Artbook, error) {

	newArtbook, err := toArtbook(input)
	if err != nil {
		return nil, err
	}

	// TODO: Here add check on existing artbook based on reference
	isExisting, err := isExistingArtbook(ctx, input.Reference)
	if isExisting {
		return nil, err
	}

	result, err := artbooksCollection.InsertOne(ctx, newArtbook)
	if err != nil {
		defer log.Println(ErrArtbookCreationFailed)
		return nil, err
	}

	newArtbook.ID = result.InsertedID.(primitive.ObjectID)

	return newArtbook, nil

}

// GetAllArtbooks finds all artbooks stored in database
func GetAllArtbooks(ctx context.Context) ([]*Artbook, error) {

	cursor, err := artbooksCollection.Find(
		ctx,
		bson.D{},
		options.Find(),
	)
	if err != nil {
		return nil, err
	}

	var artbooks []*Artbook
	if err := cursor.All(ctx, &artbooks); err != nil {
		return nil, err
	}

	return artbooks, nil

}

// GetArtbookByID finds an artbook for a given ID in database
func GetArtbookByID(ctx context.Context, id string) (*Artbook, error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	var artbook *Artbook

	err := artbooksCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&artbook)
	if err != nil {
		log.Println(ErrArtbookRetrievalFailed)
		return nil, err
	}

	return artbook, nil

}

// DeleteArtbookByID deletes an artbook for a given ID in database
func DeleteArtbookByID(ctx context.Context, id string) (deleteCount int64, err error) {

	objectID, _ := primitive.ObjectIDFromHex(id)

	result, err := artbooksCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Printf("An error occured during deletion of artbook: %s", id)
		return 0, err
	}

	if result.DeletedCount == 0 {
		log.Println(ErrArtbookRetrievalFailed)
	}

	return result.DeletedCount, nil

}

// UpdateArtbookByID finds an artbook for a given ID and update it
func UpdateArtbookByID(ctx context.Context, id string, update interface{}) (*Artbook, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var updatedArtbook Artbook

	err = artbooksCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
		options.FindOneAndUpdate().
			SetReturnDocument(options.After),
	).Decode(&updatedArtbook)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &updatedArtbook, nil

}

// toArtbook converts a create input to an artbook struct
func toArtbook(input *CreateInput) (*Artbook, error) {

	releaseDate, err := time.Parse("2006-01-02", input.ReleasedAt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	releaseDate = releaseDate.UTC()
	createdAt := time.Now().UTC()

	for _, format := range input.Formats {
		isValid := product.IsValidFormat(format)
		if !isValid {
			return nil, err
		}
	}

	_, found := product.AvailabilitiesByLabel[input.Availability]
	if !found {
		log.Printf("Availability \"%s\" is invalid", input.Availability)
		return nil, err
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
	}, nil

}

func isExistingArtbook(ctx context.Context, reference string) (bool, error) {

	result, err := artbooksCollection.FindOne(ctx, bson.M{"reference": reference}).Raw()
	if err != nil {
		return false, err
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, errors.New(ErrArtbooksAlreadyExists)

}
