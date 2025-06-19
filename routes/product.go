package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/handler"
)

func ProductRoutes(producthandler handler.ProductHandler, server *gin.Engine) {

	//category model
     server.POST("/product/addcategory",producthandler.PostProductCategory)
	 server.PUT("/product/updatecategory",producthandler.PostProductCategory)
	 server.DELETE("/product/deletecategory",producthandler.PostProductCategory)
}
