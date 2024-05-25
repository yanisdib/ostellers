package artbook

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"yanisdib/ostellers/errors"
	"yanisdib/ostellers/product"
)

type CreateInput struct {
	Reference    string                      `json:"reference" validate:"required"`
	Label        string                      `json:"label" validate:"required"`
	Description  string                      `json:"description,omitempty"`
	Categories   []string                    `json:"categories" validate:"required"`
	Tags         []string                    `json:"tags,omitempty"`
	Artists      []string                    `json:"artists" validate:"required"`
	Editors      []string                    `json:"editors" validate:"required"`
	Pictures     []product.ProductImageAttrs `json:"pictures,omitempty"`
	PagesCount   uint16                      `json:"pagesCount,omitempty"`
	Stock        uint32                      `json:"stock" validate:"required"`
	Price        float32                     `json:"price,omitempty"`
	Availability string                      `json:"availability" validate:"required"`
	Formats      []string                    `json:"formats" validate:"required"`
	ReleasedAt   string                      `json:"releasedAt,omitempty"`
}

func Create() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input *CreateInput
		defer cancel()

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := CreateArtbook(ctx, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, "Artbook created successfully")

	}

}

func GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		artbookID := c.Param("id")
		defer cancel()

		artbook, err := GetArtbookByID(ctx, artbookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, artbook)

	}

}

func DeleteByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		artbookID := c.Param("id")
		defer cancel()

		deleteCount, err := DeleteArtbookByID(ctx, artbookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		var message string
		if deleteCount == 0 {
			message = errors.ERR_ITEM_NOT_FOUND
			c.JSON(http.StatusNotFound, message)
		} else {
			message = "Artbook deleted successfully."
			c.JSON(http.StatusOK, message)
		}

	}

}

func UpdateByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		input, _ := io.ReadAll(c.Request.Body)
		defer cancel()

		update := make(map[string]interface{})

		if err := json.Unmarshal(input, &update); err != nil {
			log.Print(err)
			c.JSON(http.StatusBadRequest, "Failed to parse JSON")
			return
		}

		currentDate := time.Now().UTC()
		update["updated_at"] = currentDate

		artbookID, found := c.Params.Get("id")
		if !found {
			c.JSON(http.StatusBadRequest, "Invalid artbook ID")
			return
		}

		updatedArtbook := UpdateArtbookByID(ctx, artbookID, update)
		if updatedArtbook == nil {
			c.JSON(http.StatusBadRequest, "This artbook couldn't be updated")
			return
		}

		c.JSON(http.StatusOK, "Artbook updated successfully")

	}

}
