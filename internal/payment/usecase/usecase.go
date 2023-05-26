package usecase

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/payment"
)

type usecase struct {
	repository payment.Repository
}

func New(repository payment.Repository) payment.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetPaymentMethods() ([]models.PaymentMethod, error) {
	return u.repository.Find()
}

func (u *usecase) GetPaymentMethod(paymentID int) (models.PaymentMethod, error) {
	return u.repository.FindByID(paymentID)
}

func (u *usecase) CreatePaymentMethod(payment models.PaymentMethod) (models.PaymentMethod, error) {
	return u.repository.Create(payment)
}

func (u *usecase) UpdatePaymentMethod(paymentID int, payment models.PaymentMethod) (models.PaymentMethod, error) {
	return u.repository.UpdateByID(paymentID, payment)
}

func (u *usecase) DeletePaymentMethod(paymentID int) error {
	return u.repository.DeleteByID(paymentID)
}
