package resolvers

import (
	"github.com/tagaertner/job-stories/services/stories/generated"
	"github.com/tagaertner/job-stories/services/stories/models"
	"time"
	
)

func ToGraphQLStory(s *models.JobStory) *generated.JobStory {
	return &generated.JobStory{
		ID:        s.ID,
		UserID:    s.UserID,
		Title:     s.Title,
		Content:   s.Content,
		Tags:      s.Tags,
		Category:  s.Category,
		Mood:      *s.Mood,
		CreatedAt: s.CreatedAt.Format(time.RFC3339),
		UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
	}
}

func ToGraphQLStoryList(stories []*models.JobStory) []*generated.JobStory{
	var gqlStories []*generated.JobStory
	for _, s := range stories {
		gqlStories = append(gqlStories, ToGraphQLStory(s))
	}
	return gqlStories
}
