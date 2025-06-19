package resolvers

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"e-commerce/services/orders/generated"
	"e-commerce/services/orders/models"
	"fmt"
)


type Resolver struct {
	orders []*models.Order
}

func NewResolver() *Resolver {
	orders := []*models.Order{
		{ID: "1", UserID: "1", ProductID: "1", Quantity: 2, TotalPrice: 2599.98, Status: "completed", CreatedAt: "2025-06-15T10:30:00Z"},
		{ID: "2", UserID: "2", ProductID: "2", Quantity: 1, TotalPrice: 799.99, Status: "pending", CreatedAt: "2025-06-16T14:20:00Z"},
		{ID: "3", UserID: "1", ProductID: "3", Quantity: 1, TotalPrice: 199.99, Status: "shipped", CreatedAt: "2025-06-17T09:15:00Z"},
	}

	return &Resolver{
		orders: orders,
	}
}

// Get all orders
func (r *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	return r.Resolver.orders, nil
}

// Get one specific order
func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	for _, order := range r.orders {
		if order.ID == id {
			return order, nil
		}
	}
	return nil, fmt.Errorf("order not found")
}

// Get all orders from one user
func (r *queryResolver) OrdersByUser(ctx context.Context, userID string) ([]*models.Order, error) {
	var userOrders []*models.Order
	for _, order := range r.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct {
	orders []*models.Order
}
func NewResolver() *Resolver {
	orders := []*models.Order{
		{ID: "1", UserID: "1", ProductID: "1", Quantity: 2, TotalPrice: 2599.98, Status: "completed", CreatedAt: "2025-06-15T10:30:00Z"},
		{ID: "2", UserID: "2", ProductID: "2", Quantity: 1, TotalPrice: 799.99, Status: "pending", CreatedAt: "2025-06-16T14:20:00Z"},
		{ID: "3", UserID: "1", ProductID: "3", Quantity: 1, TotalPrice: 199.99, Status: "shipped", CreatedAt: "2025-06-17T09:15:00Z"},
	}

	return &Resolver{
		orders: orders,
	}
}
*/
