package mysql

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/product"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find() ([]models.Product, error) {
	products := []models.Product{}
	tx := r.db.Find(&products)
	return products, tx.Error
}

func (r *repository) FindByID(productID int) (models.Product, error) {
	product := models.Product{}
	tx := r.db.Where("id = ?", productID).First(&product)
	return product, tx.Error
}

func (r *repository) Create(product models.Product) (models.Product, error) {
	tx := r.db.Create(&product)
	return product, tx.Error
}

func (r *repository) UpdateByID(productID int, product models.Product) (models.Product, error) {
	tx := r.db.Where("id = ?", productID).Updates(&product)
	return product, tx.Error
}

func (r *repository) DeleteByID(productID int) error {
	tx := r.db.Where("id = ?", productID).Delete(&models.Product{})
	return tx.Error
}
