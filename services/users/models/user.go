package models

type User struct {
	ID     string  `json:"id" gorm:"primarykey"` 
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}

func (User) IsEntity() {}