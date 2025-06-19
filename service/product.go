package service

import (
	"github.com/janeefajjustin/ecommerce/models"
	"github.com/janeefajjustin/ecommerce/repo"
)

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService(productRepo *repo.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p ProductService) AddProductcategory(productcategory models.ProductCategory) error {
	err := p.repo.AddProductcategory(productcategory)
	if err != nil {
		return err
	}
	return nil
}
