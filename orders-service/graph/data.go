package graph

import "orders-service/graph/model"

// Sample orders data for testing
var orders = []*model.Order{
	{
		ID:        "1",
		UserID:    "1",  // References user from users-service
		ProductID: "1",  // References product from products-service
		Quantity:  2,
		Price:     29.99,
		Total:     59.98,
		Status:    "completed",
		CreatedAt: "2024-06-10T10:00:00Z",
	},
	{
		ID:        "2", 
		UserID:    "2",
		ProductID: "1",
		Quantity:  1,
		Price:     29.99,
		Total:     29.99,
		Status:    "pending",
		CreatedAt: "2024-06-11T14:30:00Z",
	},
	{
		ID:        "3",
		UserID:    "1", 
		ProductID: "2",
		Quantity:  3,
		Price:     15.50,
		Total:     46.50,
		Status:    "completed",
		CreatedAt: "2024-06-12T09:15:00Z",
	},
}

var nextOrderID = 4