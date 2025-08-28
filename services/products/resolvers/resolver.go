package resolvers

import (
	"github.com/tagaertner/e-commerce-graphql/services/products/services"
	"gorm.io/gorm"
)

type Resolver struct {
	ProductService *services.ProductService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		ProductService: services.NewProductService(db),
	}
}