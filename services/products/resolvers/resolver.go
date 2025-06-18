package resolvers

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"e-commerce/services/products/generated"
	"e-commerce/services/products/models"
)

type Resolver struct{}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	panic("not implemented")
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*models.Product, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
