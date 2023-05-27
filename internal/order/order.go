package order

import "github.com/mrizalr/dimy-tech-test/internal/models"

type Repository interface {
	Find() ([]models.OrderDetail, error)
	FindByID(orderID int) (models.OrderDetail, error)
	CreateOrder(orderPayload models.OrderPayload) (models.OrderDetail, error)
	UpdateOrder(orderDetailID int, orderID int, orderPayload models.OrderPayload) (models.OrderDetail, error)
	DeleteOrder(orderDetailID int, orderID int) error
}

type Usecase interface {
	GetOrders() ([]models.OrderDetail, error)
	GetOrderByID(customerID int) (models.OrderDetail, error)
	CreateOrder(order models.OrderPayload) (models.OrderDetail, error)
	UpdateOrder(orderID int, orderPayload models.OrderPayload) (models.OrderDetail, error)
	DeleteOrder(orderDetailID int) error
}
