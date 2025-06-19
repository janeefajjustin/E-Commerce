package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/models"
	"github.com/janeefajjustin/ecommerce/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler{
	return &ProductHandler{productService: productService}
}

func(p *ProductHandler) PostProductCategory(context *gin.Context){

	var productCategory models.ProductCategory
	err:=context.ShouldBindJSON(&productCategory)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}

	
	err= p.productService.AddProductcategory(productCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't add category"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
      
}
