package resolvers

import "e-commerce/services/users/services"  

type Resolver struct {
    UserService *services.UserService
}

func NewResolver() *Resolver {
    return &Resolver{
        UserService: services.NewUserService(),
    }
}