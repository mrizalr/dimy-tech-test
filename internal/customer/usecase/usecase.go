package usecase

import (
	"github.com/mrizalr/dimy-tech-test/internal/customer"
	"github.com/mrizalr/dimy-tech-test/internal/models"
)

type usecase struct {
	repository customer.Repository
}

func New(repository customer.Repository) customer.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetCustomers() ([]models.Customer, error) {
	return u.repository.Find()
}

func (u *usecase) GetCustomerByID(customerID int) (models.Customer, error) {
	return u.repository.FindByID(customerID)
}

func (u *usecase) CreateCustomer(customer models.Customer) (models.Customer, error) {
	return u.repository.Create(customer)
}

func (u *usecase) UpdateCustomer(customerID int, customer models.Customer) (models.Customer, error) {
	// Check if customer exist
	foundCustomer, err := u.repository.FindByID(customerID)
	if err != nil {
		return models.Customer{}, err
	}

	customer.ID = foundCustomer.ID
	return u.repository.UpdateByID(customerID, customer)
}

func (u *usecase) DeleteCustomer(customerID int) error {
	// Check if customer exist
	_, err := u.repository.FindByID(customerID)
	if err != nil {
		return err
	}

	return u.repository.DeleteByID(customerID)
}
