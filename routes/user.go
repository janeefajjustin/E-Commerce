package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/handler"
)

func Routes(userhandler *handler.UserHandler, server *gin.Engine) {
	// u := handler.UserHandler{}
	server.POST("/user/signup", userhandler.Signup)
	server.POST("/user/checkusername/:username", userhandler.CheckUserName)
}
