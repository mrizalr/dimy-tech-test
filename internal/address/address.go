package address

import "github.com/mrizalr/dimy-tech-test/internal/models"

type Repository interface {
	Find(customerID int) ([]models.CustomerAddress, error)
	FindByID(addressID int) (models.CustomerAddress, error)
	Create(address models.CustomerAddress) (models.CustomerAddress, error)
	UpdateByID(addressID int, address models.CustomerAddress) (models.CustomerAddress, error)
	DeleteByID(addressID int) error
}

type Usecase interface {
	GetCustomerAddresses(customerID int) ([]models.CustomerAddress, error)
	GetAddressByID(addressID int) (models.CustomerAddress, error)
	CreateAddress(address models.CustomerAddress) (models.CustomerAddress, error)
	UpdateAddress(addressID int, address models.CustomerAddress) (models.CustomerAddress, error)
	DeleteAddress(addressID int) error
}
