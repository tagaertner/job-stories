package resolvers

import (
	"github.com/tagaertner/e-commerce-graphql/services/users/services"
	"gorm.io/gorm"
)

type Resolver struct {
	UserService *services.UserService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		UserService: services.NewUserService(db),
	}
}