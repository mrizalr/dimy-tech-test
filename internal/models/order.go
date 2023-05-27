package models

type Order struct {
	ID         int              `json:"id" gorm:"primaryKey"`
	CustomerID int              `json:"customer_id" gorm:"not null"`
	Customer   *Customer        `json:"customer,omitempty"`
	AddressID  int              `json:"address_id" gorm:"not null"`
	Address    *CustomerAddress `json:"address,omitempty"`
}

type OrderDetail struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	OrderID         int            `json:"order_id" gorm:"not null"`
	Order           *Order         `json:"order,omitempty"`
	ProductID       int            `json:"product_id" gorm:"not null"`
	Product         *Product       `json:"product,omitempty"`
	Quantity        int            `json:"quantity" gorm:"not null"`
	TotalPrice      float32        `json:"total_price" gorm:"not null"`
	PaymentMethodID int            `json:"payment_method_id" gorm:"not null"`
	PaymentMethod   *PaymentMethod `json:"payment_method,omitempty"`
	IsPaid          bool           `json:"is_paid" gorm:"not null;default:false"`
}

type OrderPayload struct {
	CustomerID      int     `json:"customer_id"`
	AddressID       int     `json:"address_id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	TotalPrice      float32 `json:"total_price"`
	PaymentMethodID int     `json:"payment_method_id"`
	IsPaid          bool    `json:"is_paid"`
}
