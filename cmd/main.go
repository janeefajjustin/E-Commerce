package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/internal/db"
	"github.com/janeefajjustin/ecommerce/internal/handler/user"
	"github.com/janeefajjustin/ecommerce/internal/repo/user"
	"github.com/janeefajjustin/ecommerce/routes/user"
	"github.com/janeefajjustin/ecommerce/internal/service/user"

	
	"github.com/janeefajjustin/ecommerce/internal/handler/product"
	"github.com/janeefajjustin/ecommerce/internal/repo/product"
	"github.com/janeefajjustin/ecommerce/routes/product"
	"github.com/janeefajjustin/ecommerce/internal/service/product"
)

func main() {
	//step 1: No idea
	DB, err := db.Initialize() //step 2 : Initialize DB
	if err != nil {
		panic("db connection failed")
	}
	//step 3,4,5,6 : Initialize everything

	//user
	userrepo := user.NewUserRepo(DB)
	userserver := user.NewUserService(userrepo)
	userhandler := user.NewUserHandler(userserver)

	//product
	productrepo:=product.NewProductRepo(DB)
	productserver:=product.NewProductService(productrepo)
	producthandler:=product.NewProductHandler(productserver)

	//step 7: starting server
	server := gin.Default()

	user.UserRoutes(userhandler, server)
	product.ProductRoutes(producthandler,server)
	server.Run("localhost:8080")
	//step 8: gracefulshutdown
}
