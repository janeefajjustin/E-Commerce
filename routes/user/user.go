package user

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/internal/handler/user"
)

func UserRoutes(userhandler *user.UserHandler, server *gin.Engine) {
	// u := handler.UserHandler{}
	server.POST("/user/signup", userhandler.Signup)
	server.POST("/user/checkusername/:username", userhandler.CheckUserName)
	server.POST("/user/login", userhandler.Login)

}
