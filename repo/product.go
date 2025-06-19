package repo

import (
	"database/sql"

	"github.com/janeefajjustin/ecommerce/models"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p ProductRepo) AddProductcategory(productcategory *models.ProductCategory) error {
	return nil
}
