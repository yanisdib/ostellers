package product

type ProductRepository interface {
	CreateProduct()
	DeleteProduct()
	GetAllProducts()
	GetOneProductById()
}
