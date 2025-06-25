package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/handler"
)

func ProductRoutes(producthandler *handler.ProductHandler, server *gin.Engine) {

	//product category
     server.POST("/product/addcategory",producthandler.PostProductCategory)
	server.PUT("/product/updatecategory/:pid",producthandler.EditProductCategory)
	server.GET("/product/getcategory",producthandler.GetProductCategory)
	server.GET("/product/getcategorybyid/:pid",producthandler.GetProductCategoryByID)
    server.DELETE("/product/deletecategory/:pid",producthandler.DeleteProductCategory)


	//product image
	server.POST("/product/addimage",producthandler.PostProductImage)
	server.PUT("/product/updateimage/:pid",producthandler.EditProductImage)
	server.GET("/product/getimagebyid/:pid",producthandler.GetProductImageByID)
	server.DELETE("/product/deleteimage/:pid",producthandler.DeleteProductImage)

	//product size
	server.POST("/product/addsize",producthandler.PostProductSize)
	server.PUT("/product/updatesize/:pid",producthandler.EditProductSize)
	server.GET("/product/getsizebyid/:pid",producthandler.GetProductSizeByID)
	server.DELETE("/product/deletesize/:pid",producthandler.DeleteProductSize)

	//product
	server.POST("/product/addproduct",producthandler.PostProduct)
	// server.PUT("/product/updateproduct/:pid",producthandler.EditProduct)
	// server.GET("/product/getproductbyid/:pid",producthandler.GetProductByID)
	// server.DELETE("/product/deleteproduct/:pid",producthandler.DeleteProduct)

}
