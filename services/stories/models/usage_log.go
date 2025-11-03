package models

import "time"

type UserLog struct {
	ID 		string 	 `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID   string   `json:"userId"`
	Action      string `gorm:"type:text;not null"`
	FiltersUsed string `gorm:"type:jsonb"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}