package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/ecommerce/internal/models"
	"github.com/janeefajjustin/ecommerce/internal/service/userService"
	"github.com/janeefajjustin/ecommerce/utils"
)

type UserHandler struct {
	userService *userService.UserService
}

func NewUserHandler(userService *userService.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	
	if f:=utils.IsEmailValid(user.Email);f==false {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email id"})
		return
	}
	
	err = u.userService.SignupUser(user)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	// 	return
	// }
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can't signup the user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (u *UserHandler) CheckUserName(context *gin.Context) {

	username := context.Param("username")
	err := u.userService.CheckUserName(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "username already exists"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "username is unique"})

}


func (u *UserHandler) Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	
	if f:=utils.IsEmailValid(user.Email);f==false {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email id"})
		return
	}
	
	err = u.userService.LoginUser(user)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can't login the user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user logged in successfully"})
}