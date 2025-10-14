package service

import (
	"fmt"
	"strings"

	"github.com/sluggard/myfile/common"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/util"
)

// UserService 用户服务
// type UserService interface {
// 	Login(username string, password string) (*model.User, error)
// 	Register(user *model.User) error
// 	GetById(id uint) (*model.User, error)
// 	ResetPassword(userId uint, oldpass string, newpass string) (bool, error)
// }

// var userService = &userSer{}

// NewUserService 返回单例的用户服务
// func NewUserService() UserService {
// 	return userService
// }

// type userSer struct {
// }

var userService = &UserService{}

type UserService struct{}

func (s *UserService) Login(username string, password string) (user *model.User, err error) {
	if user, err = user.GetUserByUsername(username); err != nil {
		return nil, err
	}
	if checkPassword(user, password) {
		user.Password = "******"
		return user, nil
	}
	return nil, &common.CommonError{Message: "check password failed"}
}
func (s *UserService) Register(user *model.User) error {
	if user, _ := user.GetUserByUsername(user.Username); user.ID > 0 {
		return &common.CommonError{Message: "用户已存在"}
	}
	user.Salt = util.UUID()
	user.Password = buildPassword(user.Password, user.Salt)
	// model.DB.Create(user)

	if _, err := model.Create(user); err != nil {
		return err
	}
	libraryService.CreateLibrary(user.ID, "Default Library")
	return nil
}

func (s *UserService) GetById(id uint) (user *model.User, err error) {
	return user.GetUserById(id)
}

func (s *UserService) ResetPassword(userId uint, oldpass string, newpass string) (bool, error) {
	user, _ := s.GetById(userId)
	if checkPassword(user, oldpass) {
		user.Password = buildPassword(newpass, user.Salt)
		if _, err := model.UpdateById(user); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func checkLibraryName(name string) error {
	if "" == strings.Trim(name, " ") {
		return common.CommonError{Message: fmt.Sprintf("error Library name : '%s'", name)}
	}
	return nil
}

func checkPassword(user *model.User, password string) bool {
	return buildPassword(password, user.Salt) == user.Password
}

func buildPassword(password string, salt string) string {
	return util.Md5String(password + salt)
}
