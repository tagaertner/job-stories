package resolvers

import "github.com/tagaertner/e-commerce/services/orders/services"

type Resolver struct {
    OrderService *services.OrderService
}

func NewResolver() *Resolver {
    return &Resolver{
        OrderService: services.NewOrderService(),
    }
}