package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/routes"
)

func main() {
	server := gin.Default()
	routes.Routes(server)
	server.Run("localhost:8080")
}
