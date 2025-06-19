package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/db"
	"github.com/janeefajjustin/ecommerce/handler"
	"github.com/janeefajjustin/ecommerce/repo"
	"github.com/janeefajjustin/ecommerce/routes"
	"github.com/janeefajjustin/ecommerce/service"
)

func main() {
	//step 1: No idea
	DB, err := db.Initialize() //step 2 : Initialize DB
	if err != nil {
		panic("db connection failed")
	}
	//step 3,4,5,6 : Initialize everything

	//user
	userrepo := repo.NewUserRepo(DB)
	userserver := service.NewUserService(userrepo)
	userhandler := handler.NewUserHandler(userserver)

	//product
	productrepo:=repo.NewProductRepo(DB)
	productserver:=service.NewProductService(productrepo)
	producthandler:=handler.NewProductHandler(productserver)

	//step 7: starting server
	server := gin.Default()

	routes.UserRoutes(userhandler, server)
	routes.ProductRoutes(producthandler,server)
	server.Run("localhost:8080")
	//step 8: gracefulshutdown
}
