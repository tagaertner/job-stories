package resolvers

import (
	"context"
	"e-commerce/services/users/generated"
	"e-commerce/services/users/models"
	"fmt"
)

type Resolver struct {
	users []*models.User
}

func NewResolver() *Resolver {
	users := []*models.User{
		{ID: "1", Name: "John Doe", Email: "john@example.com", Role: "customer", Active: true},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Role: "admin", Active: true},
		{ID: "3", Name: "Bob Wilson", Email: "bob@example.com", Role: "customer", Active: false},
	}

	return &Resolver{
		users: users,
	}
}

// Get all users
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.Resolver.users, nil
}

// Get one specific user
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	for _, user := range r.Resolver.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }


