package graph

import "products-service/graph/model"

// Mock data

var (
    stock10 = int32(10)
    stock50 = int32(50)
    stock25 = int32(25)
)

var products = []*model.Product{
    {ID: "1", Name: "Laptop", Description: "Gaming laptop", Price: 999.99, Stock: &stock10, Category: "Electronics"},
    {ID: "2", Name: "Coffee Mug", Description: "Ceramic mug", Price: 15.99, Stock: &stock50, Category: "Kitchen"},
    {ID: "3", Name: "Book", Description: "Programming book", Price: 29.99, Stock: &stock25, Category: "Books"},
}
