package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/handler"
)

func Routes(server *gin.Engine) {
	server.POST("/signup", handler.Signup)
	server.POST("/checkUserName/:username",handler.CheckUserName)
}
