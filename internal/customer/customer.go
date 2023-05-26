package customer

import "github.com/mrizalr/dimy-tech-test/internal/models"

type Repository interface {
	Find() ([]models.Customer, error)
	FindByID(customerID int) (models.Customer, error)
	Create(customer models.Customer) (models.Customer, error)
	UpdateByID(customerID int, customer models.Customer) (models.Customer, error)
	DeleteByID(customerID int) error
}

type Usecase interface {
	GetCustomers() ([]models.Customer, error)
	GetCustomerByID(customerID int) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customerID int, customer models.Customer) (models.Customer, error)
	DeleteCustomer(customerID int) error
}
