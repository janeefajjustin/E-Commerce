package models

type Cart struct{
	CartID int64
	UserID int64
	TotalAmount float32
	NoOfProduct int64
}

type CartItem struct{
	ProductID int64
	Quantity int64
	CartID int64
	Price float32
}