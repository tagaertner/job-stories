package models

import (
	"time"
    "github.com/google/uuid"
    "github.com/lib/pq"
)
type JobStory struct {
    ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    UserID    string          `json:"userId"`
    Title     string          `json:"title"`
    Content   string          `json:"content"`
    Tags      pq.StringArray  `gorm:"type:text[]" json:"tags"`
    Category  string          `json:"category"`
    Mood      *string         `json:"mood,omitempty"`
    CreatedAt time.Time       `json:"createdAt"`
    UpdatedAt time.Time       `json:"updatedAt"`
    // AI-generated fields
	AITags     pq.StringArray `gorm:"type:text[];column:ai_tags" json:"aiTags"`
	AICategory *string        `gorm:"column:ai_category" json:"aiCategory,omitempty"`
	AIMood     *float64       `gorm:"column:ai_mood" json:"aiMood,omitempty"`
	AISkills   pq.StringArray `gorm:"type:text[];column:ai_skills" json:"aiSkills"`
	AIInsights *string        `gorm:"column:ai_insights" json:"aiInsights,omitempty"`
}
type PaginatedStories struct {
	Stories     []*JobStory `json:"stories"`
	TotalCount  int         `json:"totalCount"`
	CurrentPage int         `json:"currentPage"`
	HasNextPage bool        `json:"hasNextPage"`
}

func (JobStory) IsEntity() {}