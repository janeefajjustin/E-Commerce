package service

import (
	"errors"

	"github.com/janeefajjustin/ecommerce/models"
	"github.com/janeefajjustin/ecommerce/repo"
	"github.com/janeefajjustin/ecommerce/utils"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *UserService {
	return &UserService{userRepo: &userRepo}
}

func (u UserService) SignupUser(user models.User) error {

	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = u.CheckUserName(user.Username)
	if err != nil {
		return errors.New("user name already exists")
	}

	err = u.CheckPhoneNumber(user.PhoneNumber)
	if err != nil {
		return errors.New("phone number already exists")
	}

	err = u.CheckEmailID(user.Email)
	if err != nil {
		return errors.New("Email ID already exists")
	}

	err = u.userRepo.SaveUser(user)
	if err != nil {
		return errors.New("can't save user")
	}
	return nil
}

func (u UserService) CheckUserName(username string) error {
	count, err := u.userRepo.CheckUniqueUserName(username)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	return nil
}

func (u UserService) CheckPhoneNumber(phonenumber int64) error {
	count, err := u.userRepo.CheckUniquePhoneNumber(phonenumber)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("phone number already exists")
	}
	return nil
}

func (u UserService) CheckEmailID(emailId string) error {
	count, err := u.userRepo.CheckUniqueEmailID(emailId)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email id already exists")
	}
	return nil
}
