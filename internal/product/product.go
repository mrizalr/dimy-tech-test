package product

import "github.com/mrizalr/dimy-tech-test/internal/models"

type Repository interface {
	Find() ([]models.Product, error)
	FindByID(productID int) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	UpdateByID(productID int, product models.Product) (models.Product, error)
	DeleteByID(productID int) error
}

type Usecase interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(productID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(productID int, product models.Product) (models.Product, error)
	DeleteProduct(productID int) error
}
