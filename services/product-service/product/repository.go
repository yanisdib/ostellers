package product

type ProductRepository interface {
	CreateProduct(product Product) error
	DeleteProduct(id string) error
	DeleteAllProducts(ids []string) error
	GetAllProducts() []Product
	GetOneProductById() Product
}
