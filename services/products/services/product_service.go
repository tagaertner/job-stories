package services

import "e-commerce/services/products/generated"

type ProductService struct {
    products []*generated.Product
}

func NewProductService() *ProductService {
    return &ProductService{
        products: []*generated.Product{
            {
                ID: "1", 
                Name: "Laptop", 
                Price: 1299.99, 
                Description: stringPtr("High-performance laptop"), 
                Inventory: 50,
            },
            {
                ID: "2", 
                Name: "Smartphone", 
                Price: 799.99, 
                Description: stringPtr("Latest smartphone"), 
                Inventory: 100,
            },
            {
                ID: "3", 
                Name: "Headphones", 
                Price: 199.99, 
                Description: stringPtr("Wireless headphones"), 
                Inventory: 75,
            },
        },
    }
}

// Helper function to convert string to *string
func stringPtr(s string) *string {
    return &s
}

func (s *ProductService) GetAllProducts() []*generated.Product {
    return s.products
}

func (s *ProductService) GetProductByID(id string) *generated.Product {
    for _, product := range s.products {
        if product.ID == id {
            return product
        }
    }
    return nil
}