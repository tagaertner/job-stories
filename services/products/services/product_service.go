package services

import (
    "github.com/tagaertner/e-commerce-graphql/services/products/models"
    "gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	if err := s.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}