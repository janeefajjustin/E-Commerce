package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/internal/db"
	"github.com/janeefajjustin/ecommerce/internal/handler/productHandler"
	"github.com/janeefajjustin/ecommerce/internal/handler/userHandler"
	"github.com/janeefajjustin/ecommerce/internal/repo/productRepo"
	"github.com/janeefajjustin/ecommerce/internal/repo/userRepo"
	"github.com/janeefajjustin/ecommerce/internal/service/productService"
	"github.com/janeefajjustin/ecommerce/internal/service/userService"
	"github.com/janeefajjustin/ecommerce/routes/productRoutes"
	"github.com/janeefajjustin/ecommerce/routes/userRoutes"
)

func main() {
	//step 1: No idea
	DB, err := db.Initialize() //step 2 : Initialize DB
	if err != nil {
		panic("db connection failed")
	}
	//step 3,4,5,6 : Initialize everything

	//user
	userrepo := userRepo.NewUserRepo(DB)
	userserver := userService.NewUserService(userrepo)
	userhandler := userHandler.NewUserHandler(userserver)

	//product
	productrepo := productRepo.NewProductRepo(DB)
	productservice := productService.NewProductService(productrepo)
	producthandler := productHandler.NewProductHandler(productservice)

	//step 7: starting server
	server := gin.Default()

	userRoutes.UserRoutes(userhandler, server)
	productRoutes.ProductRoutes(producthandler, server)
	server.Run("localhost:8080")
	//step 8: gracefulshutdown
}
