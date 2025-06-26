package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/internal/models"
	"github.com/janeefajjustin/ecommerce/internal/service/product"
)

type ProductHandler struct {
	productService *product.ProductService
}

func NewProductHandler(productService *product.ProductService) *ProductHandler{
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
	
	context.JSON(http.StatusCreated, gin.H{"message": "new category added successfully"})
      
}

func(p *ProductHandler) EditProductCategory(context *gin.Context){

	var productCategory models.ProductCategory
	err:=context.ShouldBindJSON(&productCategory)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	
	err= p.productService.ChangeProductcategory(productCategory,pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't edit category"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "category edited successfully"})
      
}

func(p *ProductHandler) GetProductCategory(context *gin.Context){
	var productCategories []models.ProductCategory
	productCategories,err:=p.productService.GetAllProductCategory()
		if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't get category"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message":  productCategories})
      
}

func(p *ProductHandler) GetProductCategoryByID(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
    var prodCategory *models.ProductCategory
	
	prodCategory,err= p.productService.GetProductCategory(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't get category"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": *prodCategory})
      
}

func(p *ProductHandler) DeleteProductCategory(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	err=p.productService.RemoveProductCategory(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't delete category"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "category deleted"})

}


//product image

func(p *ProductHandler) PostProductImage(context *gin.Context){

	var productImage models.ProductImage
	err:=context.ShouldBindJSON(&productImage)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}

	
	err= p.productService.AddProductImage(productImage)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't add image"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "new image added successfully"})
      
}



func(p *ProductHandler) EditProductImage(context *gin.Context){

	var productImage models.ProductImage
	err:=context.ShouldBindJSON(&productImage)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	
	err= p.productService.ChangeProductImage(productImage,pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't edit image"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "image edited successfully"})
      
}

func(p *ProductHandler) GetProductImageByID(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
    var prodImage *models.ProductImage
	
	prodImage,err= p.productService.GetProductImage(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't get image"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": *prodImage})
      
}

func(p *ProductHandler) DeleteProductImage(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	err=p.productService.RemoveProductImage(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't delete image"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "image deleted"})

}

//product size

func(p *ProductHandler) PostProductSize(context *gin.Context){

	var productSize models.ProductSize
	err:=context.ShouldBindJSON(&productSize)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}

	
	err= p.productService.AddProductSize(productSize)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't add size"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "new size added successfully"})
      
}



func(p *ProductHandler) EditProductSize(context *gin.Context){

	var productSize models.ProductSize
	err:=context.ShouldBindJSON(&productSize)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
	err= p.productService.ChangeProductSize(productSize,pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't edit size"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "size edited successfully"})     
}


func(p *ProductHandler) GetProductSizeByID(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
    var prodSize *models.ProductSize
	
	prodSize,err= p.productService.GetProductSize(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't get size"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": *prodSize})
      
}


func(p *ProductHandler) DeleteProductSize(context *gin.Context){
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	err=p.productService.RemoveProductSize(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't delete size"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "size deleted"})

}



//product

func(p *ProductHandler) PostProduct(context *gin.Context){

	var product models.Product
	err:=context.ShouldBindJSON(&product)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}

	
	err= p.productService.AddProduct(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't add product"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "new product added successfully"})
      
}

func(p *ProductHandler) EditProduct(context *gin.Context){

	var product models.Product
	err:=context.ShouldBindJSON(&product)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"invalid request"} )
		return
	}
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
	err= p.productService.ChangeProduct(product,pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't edit product"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "product edited successfully"})     
}

func(p *ProductHandler) GetProductByID(context *gin.Context){

	
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}
    var prod *models.Product
	
	prod,err= p.productService.GetProduct(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't get product details"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": *prod})
      
}


func(p *ProductHandler) DeleteProduct(context *gin.Context){
	pid,err:=strconv.ParseInt(context.Param("pid"),10,64)
	if err!=nil{
		context.JSON( http.StatusBadRequest,gin.H{"message":"could not fetch the id"} )
		return
	}

	err=p.productService.RemoveProduct(pid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"can't delete product"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "product deleted"})

}

