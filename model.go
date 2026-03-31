// model.go

package main

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (p *Product) getProduct(db *gorm.DB) error {
	result := db.Limit(1).Take(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) updateProduct(db *gorm.DB) error {
	result := db.Save(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) deleteProduct(db *gorm.DB) error {
	result := db.Delete(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) createProduct(db *gorm.DB) error {
	result := db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getProducts(db *gorm.DB, start, count int) ([]Product, error) {
	var products []Product
	result := db.Offset(start).Limit(count).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
