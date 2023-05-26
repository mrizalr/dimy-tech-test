package mysql

import (
	"github.com/mrizalr/dimy-tech-test/internal/address"
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) address.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find(customerID int) ([]models.CustomerAddress, error) {
	customerAddresses := []models.CustomerAddress{}
	tx := r.db.Where("customer_id = ?", customerID).Find(&customerAddresses)
	return customerAddresses, tx.Error
}

func (r *repository) FindByID(addressID int) (models.CustomerAddress, error) {
	address := models.CustomerAddress{}
	tx := r.db.Where("id = ?", addressID).First(&address)
	return address, tx.Error
}

func (r *repository) Create(address models.CustomerAddress) (models.CustomerAddress, error) {
	tx := r.db.Create(&address)
	return address, tx.Error
}

func (r *repository) UpdateByID(addressID int, address models.CustomerAddress) (models.CustomerAddress, error) {
	tx := r.db.Where("id = ?", addressID).Updates(&address)
	return address, tx.Error
}

func (r *repository) DeleteByID(addressID int) error {
	tx := r.db.Where("id = ?", addressID).Delete(&models.CustomerAddress{})
	return tx.Error
}
