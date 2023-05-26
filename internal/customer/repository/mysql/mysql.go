package mysql

import (
	"github.com/mrizalr/dimy-tech-test/internal/customer"
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) customer.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find() ([]models.Customer, error) {
	customers := []models.Customer{}
	tx := r.db.Find(&customers)
	return customers, tx.Error
}

func (r *repository) FindByID(customerID int) (models.Customer, error) {
	customer := models.Customer{}
	tx := r.db.Where("id = ?", customerID).First(&customer)
	return customer, tx.Error
}

func (r *repository) Create(customer models.Customer) (models.Customer, error) {
	tx := r.db.Create(&customer)
	return customer, tx.Error
}

func (r *repository) UpdateByID(customerID int, customer models.Customer) (models.Customer, error) {
	tx := r.db.Where("id = ?", customerID).Updates(&customer)
	return customer, tx.Error
}

func (r *repository) DeleteByID(customerID int) error {
	tx := r.db.Where("id = ?", customerID).Delete(&models.Customer{})
	return tx.Error
}
