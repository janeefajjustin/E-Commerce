package service

import (
	"errors"

	"github.com/janeefajjustin/ecommerce/models"
	"github.com/janeefajjustin/ecommerce/repo"
)

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService(productRepo *repo.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p *ProductService) AddProductcategory(productcategory models.ProductCategory) error {
	if ! p.productRepo.IsProductCategoryValid(productcategory.CategoryName){
		return errors.New("categoryname already exists")
	}
	err := p.productRepo.AddProdcategory(productcategory)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) ChangeProductcategory(productcategory models.ProductCategory,pid int64) error {
	if ! p.productRepo.IsProductCategoryIdValid(pid){
		return errors.New("invalid product id")
	}
	err := p.productRepo.ChangeProdcategory(productcategory,pid)
	if err != nil {
		return err
	}
	return nil
}



func (p *ProductService) GetAllProductCategory() ([]models.ProductCategory,error) {
	allProductcategories,err := p.productRepo.GetAllProdCategory()
	if err != nil {
		return nil,err
	}
	return allProductcategories,err
}


func (p *ProductService) GetProductCategory(pid int64) (*models.ProductCategory,error) {
	productcategory,err := p.productRepo.GetProdCategory(pid)
	if err != nil {
		return nil,err
	}
	return productcategory,err
}

func (p *ProductService) RemoveProductCategory(pid int64) (error) {
	if ! p.productRepo.IsProductCategoryIdValid(pid){
		return errors.New("invalid product id")
	}
	err := p.productRepo.RemoveProdCategoryByID(pid)
	if err != nil {
		return err
	}
	return nil
}

//product image

func (p *ProductService) AddProductImage(productImage models.ProductImage) error {
	
	err := p.productRepo.AddProdImage(productImage)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) ChangeProductImage(productimage models.ProductImage,pid int64) error {
	if ! p.productRepo.IsProductImageIdValid(pid){
		return errors.New("invalid image id")
	}
	err := p.productRepo.ChangeProdImage(productimage,pid)
	if err != nil {
		return err
	}
	return nil
}


func (p *ProductService) GetProductImage(pid int64) (*models.ProductImage,error) {
	productimage,err := p.productRepo.GetProdImage(pid)
	if err != nil {
		return nil,err
	}
	return productimage,err
}

func (p *ProductService) RemoveProductImage(pid int64) (error) {
	if ! p.productRepo.IsProductImageIdValid(pid){
		return errors.New("invalid product image id")
	}
	err := p.productRepo.RemoveProdImageByID(pid)
	if err != nil {
		return err
	}
	return nil
}
