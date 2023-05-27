package usecase

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/order"
	"github.com/mrizalr/dimy-tech-test/internal/product"
)

type usecase struct {
	repository        order.Repository
	productRepository product.Repository
}

func New(repository order.Repository, productRepository product.Repository) order.Usecase {
	return &usecase{
		repository:        repository,
		productRepository: productRepository,
	}
}

func (u *usecase) GetOrders() ([]models.OrderDetail, error) {
	return u.repository.Find()
}

func (u *usecase) GetOrderByID(customerID int) (models.OrderDetail, error) {
	return u.repository.FindByID(customerID)
}

func (u *usecase) CreateOrder(orderPayload models.OrderPayload) (models.OrderDetail, error) {
	foundProduct, err := u.productRepository.FindByID(orderPayload.ProductID)
	if err != nil {
		return models.OrderDetail{}, err
	}

	orderPayload.TotalPrice = foundProduct.Price * float32(orderPayload.Quantity)

	order, err := u.repository.CreateOrder(orderPayload)
	if err != nil {
		return models.OrderDetail{}, err
	}

	return u.repository.FindByID(order.ID)
}

func (u *usecase) UpdateOrder(orderID int, orderPayload models.OrderPayload) (models.OrderDetail, error) {
	// Check if order exist
	foundOrder, err := u.repository.FindByID(orderID)
	if err != nil {
		return models.OrderDetail{}, err
	}

	foundProduct, err := u.productRepository.FindByID(orderPayload.ProductID)
	if err != nil {
		return models.OrderDetail{}, err
	}

	orderPayload.TotalPrice = foundProduct.Price * float32(orderPayload.Quantity)

	order, err := u.repository.UpdateOrder(orderID, foundOrder.OrderID, orderPayload)
	if err != nil {
		return models.OrderDetail{}, err
	}

	order.ID = foundOrder.ID
	return u.repository.FindByID(order.ID)
}

func (u *usecase) DeleteOrder(orderDetailID int) error {
	// Check if customer exist
	foundOrder, err := u.repository.FindByID(orderDetailID)
	if err != nil {
		return err
	}

	return u.repository.DeleteOrder(orderDetailID, foundOrder.OrderID)
}
