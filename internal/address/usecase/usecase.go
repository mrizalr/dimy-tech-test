package usecase

import (
	"github.com/mrizalr/dimy-tech-test/internal/address"
	"github.com/mrizalr/dimy-tech-test/internal/models"
)

type usecase struct {
	repository address.Repository
}

func New(repository address.Repository) address.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetCustomerAddresses(customerID int) ([]models.CustomerAddress, error) {
	return u.repository.Find(customerID)
}

func (u *usecase) GetAddressByID(addressID int) (models.CustomerAddress, error) {
	return u.repository.FindByID(addressID)
}

func (u *usecase) CreateAddress(address models.CustomerAddress) (models.CustomerAddress, error) {
	return u.repository.Create(address)
}

func (u *usecase) UpdateAddress(addressID int, address models.CustomerAddress) (models.CustomerAddress, error) {
	// Check if address exist
	foundAddress, err := u.repository.FindByID(addressID)
	if err != nil {
		return models.CustomerAddress{}, err
	}

	address.ID = foundAddress.ID
	address.CustomerID = foundAddress.CustomerID
	return u.repository.UpdateByID(addressID, address)
}

func (u *usecase) DeleteAddress(addressID int) error {
	// Check if address exist
	_, err := u.repository.FindByID(addressID)
	if err != nil {
		return err
	}

	return u.repository.DeleteByID(addressID)
}
