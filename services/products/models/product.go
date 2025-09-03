package models



type Product struct {
	 ID         string  `json:"id" gorm:"primarykey"` 
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	Inventory   int     `json:"inventory"`
}



func (Product) IsEntity() {}