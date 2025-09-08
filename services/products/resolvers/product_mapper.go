package resolvers

import (
	"github.com/tagaertner/e-commerce-graphql/services/products/generated"
	"github.com/tagaertner/e-commerce-graphql/services/products/models"
)

func ToGraphQLProduct(p *models.Product) *generated.Product {
	return &generated.Product{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		Inventory:   p.Inventory,
		Available:   p.Available,
	}
}

func ToGraphQLProductList(products []*models.Product) []*generated.Product {
	var gqlProducts []*generated.Product
	for _, p := range products {
		gqlProducts = append(gqlProducts, ToGraphQLProduct(p))
	}
	return gqlProducts
}
