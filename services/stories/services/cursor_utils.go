package services

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/tagaertner/job-stories/services/stories/models"
)

// EncodeCursor creates a cursor from a story
func EncodeCursor(story *models.JobStory) string{
	cursor := fmt.Sprintf("%s:%s", story.CreatedAt.Format(time.RFC3339Nano),story.ID.String())
	return base64.StdEncoding.EncodeToString([]byte(cursor))
}

// DecodeCursor extracts timestamp and ID from cursor
func DecodeCursor(cursor string) (time.Time, string, error){
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Time{}, "", err
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2{
		return time.Time{}, "", fmt.Errorf("infalid cursor format")
	}

	timestamp, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return time.Time{}, "", err
	}

	return timestamp, parts[1], nil
}