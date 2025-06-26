package product

import (
	"errors"
	"fmt"

	"github.com/janeefajjustin/ecommerce/internal/models"
	"github.com/janeefajjustin/ecommerce/internal/repo/product"
)

type ProductService struct {
	productRepo *product.ProductRepo
}

func NewProductService(productRepo *product.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p *ProductService) AddProductcategory(productcategory models.ProductCategory) error {

	err := p.productRepo.ProductCategoryValidation(productcategory.CategoryName)
	if err != nil {
		return errors.New("categoryname already exists")
	}
	err = p.productRepo.AddProdcategory(productcategory)
	if err != nil {

		return err
	}
	return nil
}

func (p *ProductService) ChangeProductcategory(productcategory models.ProductCategory, pid int64) error {

	err := p.productRepo.ProductCategoryIdValidation(pid)
	if err != nil {
		return err
	}
	err = p.productRepo.ChangeProdcategory(productcategory, pid)
	if err != nil {

		return err
	}
	return nil
}

func (p *ProductService) GetAllProductCategory() ([]models.ProductCategory, error) {
	allProductcategories, err := p.productRepo.GetAllProdCategory()
	if err != nil {
		//delete
		fmt.Println(err)
		return nil, err
	}
	return allProductcategories, nil
}

func (p *ProductService) GetProductCategory(pid int64) (*models.ProductCategory, error) {
	productcategory, err := p.productRepo.GetProdCategory(pid)
	if err != nil {
		return nil, err
	}
	return productcategory, nil
}

func (p *ProductService) RemoveProductCategory(pid int64) error {
	err := p.productRepo.ProductCategoryIdValidation(pid)
	if err != nil {
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

func (p *ProductService) ChangeProductImage(productimage models.ProductImage, pid int64) error {
	if !p.productRepo.IsProductImageIdValid(pid) {
		return errors.New("invalid image id")
	}
	err := p.productRepo.ChangeProdImage(productimage, pid)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetProductImage(pid int64) (*models.ProductImage, error) {
	productimage, err := p.productRepo.GetProdImage(pid)
	if err != nil {
		return nil, err
	}
	return productimage, err
}

func (p *ProductService) RemoveProductImage(pid int64) error {
	if !p.productRepo.IsProductImageIdValid(pid) {
		return errors.New("invalid product image id")
	}
	err := p.productRepo.RemoveProdImageByID(pid)
	if err != nil {
		return err
	}
	return nil
}

//product size

func (p *ProductService) AddProductSize(productSize models.ProductSize) error {

	err := p.productRepo.ProductSizeValidation(productSize.Size, productSize.ProductID)
	if err != nil {
		return errors.New("sizename already exists for this product")
	}

	err = p.productRepo.AddProdSize(productSize)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) ChangeProductSize(productsize models.ProductSize, pid int64) error {
	err := p.productRepo.ProductSizeIdValidation(pid)
	if err != nil {
		return err
	}
	err = p.productRepo.ChangeProdSize(productsize, pid)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetProductSize(pid int64) (*models.ProductSize, error) {

	err := p.productRepo.ProductSizeIdValidation(pid)
	if err != nil {
		return nil, err
	}

	productsize, err := p.productRepo.GetProdSize(pid)
	if err != nil {
		return nil, err
	}
	return productsize, err
}

func (p *ProductService) RemoveProductSize(pid int64) error {
	err := p.productRepo.ProductSizeIdValidation(pid)
	if err != nil {
		return err
	}
	err = p.productRepo.RemoveProdSizeByID(pid)
	if err != nil {
		return err
	}
	return nil
}

//product

func (p *ProductService) AddProduct(product models.Product) error {

	err := p.productRepo.ProductValidation(product.ProductName)
	if err != nil {
		return errors.New("product already exists")
	}

	err = p.productRepo.AddProd(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) ChangeProduct(product models.Product, pid int64) error {
	err := p.productRepo.ProductIdValidation(pid)
	if err != nil {
		return err
	}
	err = p.productRepo.ChangeProd(product, pid)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetProduct(pid int64) (*models.Product, error) {

	err := p.productRepo.ProductIdValidation(pid)
	if err != nil {
		return nil, err
	}

	product, err := p.productRepo.GetProd(pid)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (p *ProductService) RemoveProduct(pid int64) error {
	err := p.productRepo.ProductIdValidation(pid)
	if err != nil {
		return err
	}
	err = p.productRepo.RemoveProd(pid)
	if err != nil {
		return err
	}
	return nil
}
