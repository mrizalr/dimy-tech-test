package mysql

import (
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/order"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find() ([]models.OrderDetail, error) {
	orders := []models.OrderDetail{}
	tx := r.db.Preload("Product").Preload("Order").Preload("Order.Customer").Preload("Order.Address").Find(&orders)
	return orders, tx.Error
}

func (r *repository) FindByID(orderID int) (models.OrderDetail, error) {
	order := models.OrderDetail{}
	tx := r.db.Preload("Product").Preload("Order").Preload("Order.Customer").Preload("Order.Address").Where("id = ?", orderID).First(&order)
	return order, tx.Error
}

func (r *repository) CreateOrder(orderPayload models.OrderPayload) (models.OrderDetail, error) {
	order := models.Order{
		CustomerID: orderPayload.CustomerID,
		AddressID:  orderPayload.AddressID,
	}

	err := r.db.Create(&order).Error
	if err != nil {
		return models.OrderDetail{}, err
	}

	orderDetail := models.OrderDetail{
		OrderID:         order.ID,
		ProductID:       orderPayload.ProductID,
		Quantity:        orderPayload.Quantity,
		TotalPrice:      orderPayload.TotalPrice,
		PaymentMethodID: orderPayload.PaymentMethodID,
		IsPaid:          orderPayload.IsPaid,
	}

	err = r.db.Create(&orderDetail).Error
	if err != nil {
		return models.OrderDetail{}, err
	}

	return orderDetail, nil
}

func (r *repository) UpdateOrder(orderDetailID int, orderID int, orderPayload models.OrderPayload) (models.OrderDetail, error) {
	order := models.Order{
		AddressID: orderPayload.AddressID,
	}

	err := r.db.Where("id = ?", orderID).Updates(&order).Error
	if err != nil {
		return models.OrderDetail{}, err
	}

	orderDetail := models.OrderDetail{
		ProductID:       orderPayload.ProductID,
		Quantity:        orderPayload.Quantity,
		TotalPrice:      orderPayload.TotalPrice,
		PaymentMethodID: orderPayload.PaymentMethodID,
		IsPaid:          orderPayload.IsPaid,
	}

	err = r.db.Where("id = ?", orderDetailID).Updates(&orderDetail).Error
	if err != nil {
		return models.OrderDetail{}, err
	}

	return orderDetail, nil
}

func (r *repository) DeleteOrder(orderDetailID int, orderID int) error {
	err := r.db.Where("id = ?", orderDetailID).Delete(&models.OrderDetail{}).Error
	if err != nil {
		return err
	}

	err = r.db.Where("id = ?", orderID).Delete(&models.Order{}).Error
	if err != nil {
		return err
	}

	return nil
}
