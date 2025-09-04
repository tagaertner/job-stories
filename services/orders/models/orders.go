package models

type Order struct {
	 ID        string  `json:"id" gorm:"primarykey"`  
	UserID     string  `json:"userId"`
	ProductID  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
	Status     string  `json:"status"`
	CreatedAt  Time    `json:"createdAt"` // <- custom Time type
}

func (Order) IsEntity() {}





