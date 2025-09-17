package services

import (
	"github.com/tagaertner/job-stories/services/stories/models"
	"gorm.io/gorm"
	"time"
	"fmt"
	"context"
	"errors"
)

type StoryService struct {
	db *gorm.DB
}

func NewStoryService(db *gorm.DB) *StoryService {
	return &StoryService{db: db}
}

func (s *StoryService) GetAllStories() ([]*models.JobStory, error) {
	var stories []*models.JobStory
	if err := s.db.Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

func (s *StoryService) GetStoryByID(id string) (*models.JobStory, error) {
	var story models.JobStory
	if err := s.db.First(&story, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &story, nil
}

func (s *StoryService) GetStoriesByUser(ctx context.Context, userID string) ([]*models.JobStory, error) {
    var stories []*models.JobStory
    if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Find(&stories).Error; err != nil {
        return nil, err
    }
    return stories, nil
}



func (s *StoryService) CreateStory(
	ctx context.Context,  
	input models.CreateStoryInput,
) (*models.JobStory, error) {
	story := &models.JobStory{
		ID:        fmt.Sprintf("story_%d", time.Now().UnixNano()),
		UserID:    input.UserID,
		Title:     input.Title,
		Content:   input.Content,
		Tags:      input.Tags,
		Category:  input.Category,
		Mood:      input.Mood,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}


	if err := s.db.WithContext(ctx).Create(story).Error; err != nil {
		return nil, err
	}
	return story, nil
}

func (s *StoryService)UpdateStory(
	ctx context.Context,
	id string, 
	input models.UpdateStoryInput)(*models.JobStory, error){

	story := &models.JobStory{ID: id}

	updates := s.db.WithContext(ctx).Model(&story)

	if input.Title != nil{
		updates = updates.Update("title", *input.Title)
	}
	
	if input.Content != nil{
		updates = updates.Update("content", *input.Content)
	}
	if input.Category != nil{
		updates = updates.Update("category", *input.Category)
	}
	if input.Mood != nil{
		updates = updates.Update("mood", *input.Mood)
	}

	// Execute the update
	if err := updates.Error; err != nil{
		return nil, err
	}

	// Return the updated user
	if err := s.db.WithContext(ctx).First(&story, "id = ?", id).Error; err != nil{
		return nil, err
	}
	return story, nil
}

func (s *StoryService)DeleteStory(
	ctx context.Context, 
	input *models.DeleteStoryInput) (bool, error){
	var result *gorm.DB

	if input.ID == nil && input.Title == nil {
		return false, errors.New("either ID or Title must be provided for deletion")
	}

	if input.ID != nil {
		result = s.db.WithContext(ctx).Delete(&models.JobStory{}, "id = ?", input.ID)
	} else if input.Title != nil {
		result = s.db.WithContext(ctx).Delete(&models.JobStory{}, "title = ?", input.Title)
	}

	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}





