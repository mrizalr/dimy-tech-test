package models

type Product struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"not null"`
	Price float32 `json:"price" gorm:"not null"`
}
