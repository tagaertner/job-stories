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

type CreateOrderInput struct {
	UserID     string  `json:"userId"`
	ProductID  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
	Status     string  `json:"status"`
	CreatedAt  Time    `json:"createdAt"` // <- custom Time type
}

type UpdateOrderInput struct {
	OrderID     string   `json:"orderId"`
	Quantity    *int     `json:"quantity"`
	TotalPrice  *float64 `json:"totalPrice"`
	Status      *string  `json:"status"`
}

type DeleteOrderInput struct {
	OrderID     string   `json:"orderId"`
	UserID	    string   `json:"userId"`
}

type SetOrderStatusInput struct {
	OrderID     string   `json:"orderId"`
	Status      *string  `json:"status"`
}

type ChangeOrderQuantityInput struct {
	OrderID     string   `json:"orderId"`
	Quantity   int     `json:"quantity"`
}

func (Order) IsEntity() {}





