package artbook

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

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

type GetArtbooksOutput struct {
	Count    int        `json:"count"`
	Previous string     `json:"previous,omitempty"`
	Next     string     `json:"next,omitempty"`
	Results  []*Artbook `json:"results"`
}

func Create() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input *CreateInput
		defer cancel()

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to parse JSON",
			})

			return
		}

		newArtbook, err := CreateArtbook(ctx, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": "Failed to create a new artbook",
				"detail":  "Ensure that the ID provided in the request is correct",
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Artbook created successfully",
			"result":  newArtbook,
		})

	}

}

func GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		artbooks, err := GetAllArtbooks(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": "Failed to retrieve artbooks",
				"detail":  "<customMessage>",
			})
			return
		}

		if artbooks == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "No artbooks found",
				"results": artbooks,
			})
		} else {
			c.JSON(
				http.StatusOK,
				GetArtbooksOutput{
					Count:   len(artbooks),
					Next:    "localhost:6060/artbooks?limit=10&offset=10",
					Results: artbooks,
				},
			)
		}

	}

}

func GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		artbookID := c.Param("id")
		defer cancel()

		artbook, err := GetArtbookByID(ctx, artbookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": "Failed to retrieve this artbook",
				"detail":  "Ensure that the ID provided in the request is correct",
			})

			return
		}

		c.JSON(http.StatusOK, artbook)

	}

}

func DeleteByID() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		artbookID, found := c.Params.Get("id")
		defer cancel()

		if !found {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Missing ID parameter",
				"detail":  "Ensure that an ID is provided to the request",
			})

			return
		}

		deleteCount, err := DeleteArtbookByID(ctx, artbookID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to delete this artbook",
				"detail":  "Ensure that the ID provided in the request is correct",
			})

			return
		}

		if deleteCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Artbook not found",
				"detail":  "Ensure that the ID provided in the request is correct",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count":   deleteCount,
				"status":  http.StatusOK,
				"message": "Artbook deleted successfully",
			})
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
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to parse JSON",
			})

			return
		}

		currentDate := time.Now().UTC()
		update["updated_at"] = currentDate

		artbookID, found := c.Params.Get("id")
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Missing ID parameter",
				"detail":  "Ensure that an ID is provided to the request",
			})

			return
		}

		updatedArtbook, err := UpdateArtbookByID(ctx, artbookID, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to update this artbook",
				"detail":  "Ensure that the ID and updated data provided in the request are correct",
			})

			return
		}

		if updatedArtbook == nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": "Artbook not found",
					"detail":  "Ensure that the ID provided in the request is correct",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Artbook updated successfully",
				"result":  updatedArtbook,
			})
		}

	}

}
