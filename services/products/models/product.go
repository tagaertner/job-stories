package models



type Product struct {
	 ID         string  `json:"id" gorm:"primarykey"` 
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	Inventory   int     `json:"inventory"`
	Available   bool    `json:"available"`
}

type CreateProductInput struct {
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description *string  `json:"description"`
	Inventory   int      `json:"inventory"`
}

type UpdateProductInput struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
	Inventory   *int     `json:"inventory"`
}

type DeleteProductInput struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}


func (Product) IsEntity() {}