package models

type PaymentMethod struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	IsActive bool   `json:"is_active" gorm:"not null;default:false"`
}
