package artbook

import (
	"context"
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
