package services

import (
	"github.com/tagaertner/job-stories/services/stories/models"
	"gorm.io/gorm"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tagaertner/job-stories/services/stories/generated"
	
	"context"
	"errors"
)

type StoryService struct {
	db *gorm.DB
}

func NewStoryService(db *gorm.DB) *StoryService {
	return &StoryService{db: db}
}

// GetAllStories returns filtered stories from DB
func (s *StoryService) GetAllStories(filter *generated.StoryFilter, limit *int, offset *int) ([]*models.JobStory, error) {
	var stories []*models.JobStory

	// Default pagination
	defaultLimit := 100
	defaultOffset := 0
	if limit != nil {
		defaultLimit = *limit
	}
	if offset != nil {
		defaultOffset = *offset
	}

	// Start query
	query := s.db.Model(&models.JobStory{})

	// Apply filters
	if filter != nil {
		if filter.Category != nil {
			query = query.Where("category = ?", *filter.Category)
		}
		if filter.Mood != nil {
			query = query.Where("mood = ?", *filter.Mood)
		}
		if len(filter.Tags) > 0 {
			query = query.Where("tags && ?", pq.Array(filter.Tags))
		}
		if filter.SearchText != nil {
			searchPattern := "%" + *filter.SearchText + "%"
			query = query.Where("title ILIKE ? OR content ILIKE ?", searchPattern, searchPattern)
		}
		if filter.DateFrom != nil {
			query = query.Where("created_at >= ?", *filter.DateFrom)
		}
		if filter.DateTo != nil {
			query = query.Where("created_at <= ?", *filter.DateTo)
		}
	}

	// Apply pagination + execute
	if err := query.
		Limit(defaultLimit).
		Offset(defaultOffset).
		Order("created_at DESC").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}


func (s *StoryService) GetStoriesByUser(ctx context.Context, userID string, page int, pageSize int) ([]*models.JobStory, int, error) {
	var stories []*models.JobStory
	var total int64

	offset := (page - 1) * pageSize

	// Get total count
	if err := s.db.Model(&models.JobStory{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	if err := s.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&stories).Error; err != nil {
		return nil, 0, err
	}

	return stories, int(total), nil
}

func (s *StoryService)GetStoryByID(ctx context.Context, id string)(*models.JobStory, error){
	var story models.JobStory 
	if err := s.db.First(&story, "id = ?", id).Error; err != nil{
		return nil, err
	}
	return &story, nil
}

func (s *StoryService) GetStoriesByUserCursor(ctx context.Context, userID string, after *string, first int) ([]*models.JobStory, bool, error) {
	var stories []*models.JobStory
	
	query := s.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC, id DESC")
	
	// If we have a cursor, decode it and filter
	if after != nil && *after != "" {
		timestamp, id, err := DecodeCursor(*after)
		if err != nil {
			return nil, false, fmt.Errorf("invalid cursor: %w", err)
		}
		
		// Filter for stories before this cursor
		query = query.Where(
			"(created_at < ?) OR (created_at = ? AND id < ?)",
			timestamp, timestamp, id,
		)
	}
	
	// Fetch first + 1 to check if there's a next page
	if err := query.Limit(first + 1).Find(&stories).Error; err != nil {
		return nil, false, err
	}
	
	// Check if there are more results
	hasNextPage := len(stories) > first
	if hasNextPage {
		stories = stories[:first]
	}
	
	return stories, hasNextPage, nil
}


func (s *StoryService) CreateStory(
	ctx context.Context,  
	input models.CreateStoryInput,
) (*models.JobStory, error) {
	story := &models.JobStory{
		ID:        uuid.New(),
		UserID:    input.UserID,
		Title:     input.Title,
		Content:   input.Content,
		Tags:      pq.StringArray(input.Tags),
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

	story := &models.JobStory{}

	updates := s.db.WithContext(ctx).Model(&story).Where("id= ?", id)

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





