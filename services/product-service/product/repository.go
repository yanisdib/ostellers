package product

import (
	"context"
	"service/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, ids []string) error
	GetProducts(ctx context.Context) []*models.Product
	GetOneById(ctx context.Context) *models.Product
}
