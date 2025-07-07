package services

import "github.com/tagaertner/e-commerce/services/orders/generated"

type OrderService struct {
    orders []*generated.Order
    users  []*generated.User
}

func NewOrderService() *OrderService {
    return &OrderService{
        orders: []*generated.Order{
            {ID: "1", UserID: "1", ProductID: "1", Quantity: 2, TotalPrice: 2599.98, Status: "completed", CreatedAt: "2025-06-15T10:30:00Z"},
            {ID: "2", UserID: "2", ProductID: "2", Quantity: 1, TotalPrice: 799.99, Status: "pending", CreatedAt: "2025-06-16T14:20:00Z"},
            {ID: "3", UserID: "1", ProductID: "3", Quantity: 1, TotalPrice: 199.99, Status: "shipped", CreatedAt: "2025-06-17T09:15:00Z"},
        },
        users: []*generated.User{
            // Add users if needed
        },
    }
}

func (s *OrderService) GetAllOrders() []*generated.Order {
    return s.orders
}

func (s *OrderService) GetOrderByID(id string) *generated.Order {
    for _, order := range s.orders {
        if order.ID == id {
            return order
        }
    }
    return nil
}

func (s *OrderService) GetOrdersByUserID(userID string) []*generated.Order {
    var userOrders []*generated.Order
    for _, order := range s.orders {
        if order.UserID == userID {
            userOrders = append(userOrders, order)
        }
    }
    return userOrders
}

func (s *OrderService) GetUserByID(id string) *generated.User {
    for _, user := range s.users {
        if user.ID == id {
            return user
        }
    }
    return nil
}