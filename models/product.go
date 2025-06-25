package models

import "time"


type ProductCategory struct{
	CategoryID int64
	CategoryName string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ProductImage struct{
	ProductImageID int64
	ProductID int64
	ImageURL string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ProductSize struct{
    SizeID int64
	Size string
	ProductID int64
	ProductQuantity int64
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type Product struct{
	ProductName string
	ProductID int64
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

}