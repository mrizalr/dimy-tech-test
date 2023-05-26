package mysql

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/payment"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find() ([]models.PaymentMethod, error) {
	paymentMethods := []models.PaymentMethod{}
	tx := r.db.Find(&paymentMethods)
	return paymentMethods, tx.Error
}

func (r *repository) FindByID(paymentMethodID int) (models.PaymentMethod, error) {
	paymentMethod := models.PaymentMethod{}
	tx := r.db.Where("id = ?", paymentMethodID).First(&paymentMethod)
	return paymentMethod, tx.Error
}

func (r *repository) Create(paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	tx := r.db.Create(&paymentMethod)
	return paymentMethod, tx.Error
}

func (r *repository) UpdateByID(paymentMethodID int, paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	tx := r.db.Where("id = ?", paymentMethodID).Updates(&paymentMethod)
	return paymentMethod, tx.Error
}

func (r *repository) DeleteByID(paymentMethodID int) error {
	tx := r.db.Where("id = ?", paymentMethodID).Delete(&models.PaymentMethod{})
	return tx.Error
}
