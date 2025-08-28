package models

type User struct {
	ID string
}

func (User) IsEntity() {}