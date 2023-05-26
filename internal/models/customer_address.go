package models

type CustomerAddress struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	CustomerID int    `json:"customer_id" gorm:"not null;"`
	Address    string `json:"address" gorm:"not null"`
	Customer   Customer
}
