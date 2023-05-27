package payment

import "github.com/mrizalr/dimy-tech-test/internal/models"

type Repository interface {
	Find() ([]models.PaymentMethod, error)
	FindByID(paymentMethodID int) (models.PaymentMethod, error)
	Create(paymentMethod models.PaymentMethod) (models.PaymentMethod, error)
	UpdateByID(paymentMethodID int, paymentMethod models.PaymentMethod) (models.PaymentMethod, error)
	DeleteByID(paymentMethodID int) error
}

type Usecase interface {
	GetPaymentMethods() ([]models.PaymentMethod, error)
	GetPaymentMethodByID(paymentID int) (models.PaymentMethod, error)
	CreatePaymentMethod(payment models.PaymentMethod) (models.PaymentMethod, error)
	UpdatePaymentMethod(paymentID int, payment models.PaymentMethod) (models.PaymentMethod, error)
	DeletePaymentMethod(paymentID int) error
}
