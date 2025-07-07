package resolvers

import "e-commerce/services/products/services"

type Resolver struct {
    ProductService *services.ProductService
}

func NewResolver() *Resolver {
    return &Resolver{
        ProductService: services.NewProductService(),
    }
}