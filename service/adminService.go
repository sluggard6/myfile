package service

import (
	"github.com/sluggard/myfile/dao"
	"github.com/sluggard/myfile/model"
)

type UserService interface {
	Login(user model.User) (model.User, error)
}

func NewUserService() UserService {
	return userService{dao.NewUserDao()}
}

type userService struct {
	// repo repositories.MovieRepository
	userDao *dao.UserDao
}

func (s *userService) Login(user model.User) {

}
