package models

import (
	"errors"
	"retail-app/config"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}

func CreateOrder(customerID uint, items []OrderItem) (*Order, error) {
	var lastOrder Order
	result := config.DB.Where("customer_id=?", customerID).Order("created_at DESC").First(&lastOrder)
	if result.Error == nil && time.Since(lastOrder.CreatedAt) < 5*time.Minute {
		return nil, errors.New("cannot place multiple orders within 5 minutes")
	}

	tx := config.DB.Begin()
	order := Order{CustomerID: customerID}
	for _, item := range items {
		var product Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if product.Quantity < item.Quantity {
			return nil, errors.New("insufficient stock for product ID ")
		}
		product.Quantity -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		item.Price = product.Price
		order.Items = append(order.Items, item)
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &order, nil

}

func GetOrdersByCustomer(customerID uint) ([]Order, error) {
	var orders []Order
	err := config.DB.Preload("Items").Where("customer_id = ?", customerID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func GetAllOrders() ([]Order, error) {
	var orders []Order
	err := config.DB.Preload("Items").Order("created_at desc").Find(&orders).Error
	return orders, err
}
