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
}
