package services

import (
	"context"
	"fmt"
	"time"

	"github.com/tagaertner/e-commerce-graphql/services/users/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// Query
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Mutation
func (s *UserService)CreateUser(ctx context.Context, name, email string) (*models.User, error){
	user := &models.User{
		ID: fmt.Sprintf("user_%d", time.Now().UnixNano()),
		Name:   name,
		Email:  email,
		Role:   "user",   // default role
		Active: true, 
	} 
	if err := s.db.WithContext(ctx).Create(user).Error; err !=nil {
		return nil, err
	}
	return user, nil

}

func (s *UserService)UpdateUser(ctx context.Context,id string, name, email, role *string, active *bool) (*models.User, error){
	user := &models.User{ID: id}

	updates := s.db.WithContext(ctx).Model(&user)

	if name != nil{
		updates = updates.Update("name", *name)
	}
	if email != nil {
		updates = updates.Update("email", *email)
	}
	if role != nil{
		updates = updates.Update("role", *role)
	}
	if active != nil{
		updates = updates.Update("active", *active)
	}

	// Execute the update
	if err := updates.Error; err != nil{
		return nil, err
	}

	// Return the updated user
	if err := s.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil{
		return nil, err
	}
	return user, nil
}

func (s *UserService)DeleteUser(ctx context.Context, input models.DeleteUserInput) (bool, error){
	var result *gorm.DB

	if input.ID != nil{
		result = s.db.WithContext(ctx).Delete(&models.User{}, "id = ?", input.ID)
	} else if input.Name != nil {
		result = s.db.WithContext(ctx).Delete(&models.User{}, "name = ?", input.Name)
	}

	if result.Error != nil{
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

