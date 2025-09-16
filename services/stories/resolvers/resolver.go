package resolvers

import (
	"github.com/tagaertner/job-stories/services/stories/services"
	"gorm.io/gorm"
)

type Resolver struct {
	StoryService *services.StoryService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		StoryService: services.NewStoryService(db),
	}
}