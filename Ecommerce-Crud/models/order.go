package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}
