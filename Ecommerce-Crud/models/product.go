package models

import (
	"retail-app/config"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func CreateProduct(product *Product) error {
	return config.DB.Create(product).Error
}

func GetProducts() ([]Product, error) {
	var products []Product
	err := config.DB.Find(&products).Error
	return products, err
}

func UpdateProduct(id uint, updated Product) (*Product, error) {
	var product Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	if updated.Name != "" {
		product.Name = updated.Name
	}
	if updated.Description != "" {
		product.Description = updated.Description
	}
	if updated.Price != 0 {
		product.Price = updated.Price
	}
	if updated.Quantity != 0 {
		product.Quantity = updated.Quantity
	}

	config.DB.Save(&product)
	return &product, nil

}
