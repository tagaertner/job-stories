package graph

import "users-service/graph/model"

// dummy data

var users = []*model.User{
    {ID: "1", Name: "John Doe", Email: "john@example.com", Password: "hashed123", CreatedAt: "2024-01-01", IsActive: true},
    {ID: "2", Name: "Jane Smith", Email: "jane@example.com", Password: "hashed456", CreatedAt: "2024-01-02", IsActive: true},
}