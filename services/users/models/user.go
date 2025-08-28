package models

type User struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}

func (User) IsEntity() {}