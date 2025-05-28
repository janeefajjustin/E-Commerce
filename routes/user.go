package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/handler"
)

func Routes(server *gin.Engine) {
	u := handler.UserHandler{}
	server.POST("/user/signup", u.Signup)
	server.POST("/user/checkusername/:username", u.CheckUserName)
}
