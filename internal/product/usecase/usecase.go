package usecase

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/product"
)

type usecase struct {
	repository product.Repository
}

func New(repository product.Repository) product.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetProducts() ([]models.Product, error) {
	return u.repository.Find()
}

func (u *usecase) GetProductByID(productID int) (models.Product, error) {
	return u.repository.FindByID(productID)
}

func (u *usecase) CreateProduct(product models.Product) (models.Product, error) {
	return u.repository.Create(product)
}

func (u *usecase) UpdateProduct(productID int, product models.Product) (models.Product, error) {
	// Check if product exist
	foundProduct, err := u.repository.FindByID(productID)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = foundProduct.ID
	return u.repository.UpdateByID(productID, product)
}

func (u *usecase) DeleteProduct(productID int) error {
	// Check if product exist
	_, err := u.repository.FindByID(productID)
	if err != nil {
		return err
	}

	return u.repository.DeleteByID(productID)
}
