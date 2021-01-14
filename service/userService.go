package service

import (
	"github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/model"
)

type UserService interface {
	Login(username string, password string) (*model.User, string)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
	// repo repositories.MovieRepository
	// userDao dao.UserDao
}

func (s *userService) Login(username string, password string) (*model.User, string) {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err.Error()
	}
	if checkPassword(user, password) {
		return user, "success"
	}
	return nil, "check password failed"
}
func checkPassword(user *model.User, password string) bool {
	return user.Password == password
}
