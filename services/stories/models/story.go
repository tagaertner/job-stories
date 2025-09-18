package models

import (
	"time"
    "gorm.io/gorm"
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
}

func (s *JobStory) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

func (JobStory) IsEntity() {}