package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/models"
	"github.com/janeefajjustin/ecommerce/service"
)

type UserHandler struct{
	userService service.UserService
}

func NewUserHandler(userService service.UserService)UserHandler{
	return UserHandler{userService: userService}
}

func(u UserHandler) Signup(context *gin.Context){
	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest ,gin.H{"message":"Invalid Request"})	
		return
	}
	err=u.userService.SignupUser(user)
	if err!=nil{
		context.JSON(http.StatusInternalServerError ,gin.H{"message":"can't signup the user"})	
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"user created successfully"})
}

func(u UserHandler) CheckUserName(context *gin.Context){
	
    username:=context.Param("username")
	err:=u.userService.CheckUserName(username)
	if err!=nil{
		context.JSON(http.StatusBadRequest ,gin.H{"message":"username already exists"})	
		return
	}
	context.JSON(http.StatusOK, gin.H{"message":"username is unique"})

}
