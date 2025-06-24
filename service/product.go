package service

import (
	"errors"
	"fmt"

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
	fmt.Println(productcategory.CategoryName)
	err:=p.productRepo.ProductCategoryValidation(productcategory.CategoryName)
	if  err!=nil{
		return errors.New("categoryname already exists")	
	}
	err = p.productRepo.AddProdcategory(productcategory)
	if err != nil {

		return err
	}
	return nil
}

func (p *ProductService) ChangeProductcategory(productcategory models.ProductCategory,pid int64) error {
	
	err:=p.productRepo.ProductCategoryIdValidation(pid)
	if err!=nil{
		return err
	}
	err = p.productRepo.ChangeProdcategory(productcategory,pid)
	if err != nil {
		
		return err
	}
	return nil
}



func (p *ProductService) GetAllProductCategory() ([]models.ProductCategory,error) {
	allProductcategories,err := p.productRepo.GetAllProdCategory()
	if err != nil {
		//delete
		fmt.Println(err)
		return nil,err
	}
	return allProductcategories,nil
}


func (p *ProductService) GetProductCategory(pid int64) (*models.ProductCategory,error) {
	productcategory,err := p.productRepo.GetProdCategory(pid)
	if err != nil {
		return nil,err
	}
	return productcategory,nil
}

func (p *ProductService) RemoveProductCategory(pid int64) (error) {
	err:=p.productRepo.ProductCategoryIdValidation(pid)
	if err!=nil{
		return err
	}
	err = p.productRepo.RemoveProdCategoryByID(pid)
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
