package resolvers

import (
    "context"
    "e-commerce/services/products/models"
    "fmt"
)

type Resolver struct {
    products []*models.Product
}

func NewResolver() *Resolver {
    products := []*models.Product{
        {ID: "1", Name: "Gaming Laptop", Price: 1299.99, Description: stringPtr("High-performance gaming laptop"), Inventory: 15},
        {ID: "2", Name: "Smartphone", Price: 799.99, Description: stringPtr("Latest flagship smartphone"), Inventory: 50},
        {ID: "3", Name: "Wireless Headphones", Price: 199.99, Description: stringPtr("Noise-canceling wireless headphones"), Inventory: 30},
    }

    return &Resolver{
        products: products,
    }
}

func (r *Resolver) Products(ctx context.Context) ([]*models.Product, error) {
    return r.products, nil
}

func (r *Resolver) Product(ctx context.Context, id string) (*models.Product, error) {
    for _, product := range r.products {
        if product.ID == id {
            return product, nil
        }
    }
    return nil, fmt.Errorf("product not found")
}

func stringPtr(s string) *string {
    return &s
}
