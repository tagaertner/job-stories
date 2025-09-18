package models

import (
	"time"
)
type JobStory struct {
    ID        string    `json:"id" gorm:"primaryKey"`
    UserID    string    `json:"userId" gorm:"not null"`
    Title     string    `json:"title" gorm:"not null"`
    Content   string    `json:"content" gorm:"type:text;not null"`
    Tags      []string  `gorm:"type:text[]"`
    Category  string    `json:"category"` // "achievement", "learning", "blocker", "collaboration"
    Mood      *string    `json:"mood"` // "positive", "neutral", "negative"
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
type PaginatedStories struct {
	Stories     []*JobStory `json:"stories"`
	TotalCount  int         `json:"totalCount"`
	CurrentPage int         `json:"currentPage"`
	HasNextPage bool        `json:"hasNextPage"`
}

func (JobStory) IsEntity() {}