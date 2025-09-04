package models

type User struct {
	ID string `json:"id"`
	
}

func (User) IsEntity() {}