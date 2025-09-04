package models

type Product struct {
	ID string `json:"id"`
}

func (Product) IsEntity() {}