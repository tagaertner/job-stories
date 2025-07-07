package services

import "e-commerce/services/users/generated"  // ✅ Correct local path

type UserService struct {
    users []*generated.User
}

func NewUserService() *UserService {
    return &UserService{
        users: []*generated.User{
            {ID: "1", Name: "Alice", Email: "alice@example.com", Role: "customer", Active: true},
            {ID: "2", Name: "Bob", Email: "bob@example.com", Role: "admin", Active: true},
            {ID: "3", Name: "Charlie", Email: "charlie@example.com", Role: "customer", Active: false},
        },
    }
}

func (s *UserService) GetAllUsers() []*generated.User {
    return s.users
}

func (s *UserService) GetUserByID(id string) *generated.User {
    for _, user := range s.users {
        if user.ID == id {
            return user
        }
    }
    return nil
}