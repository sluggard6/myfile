package service

import (
	"github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/common"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/util"
)

type UserService interface {
	Login(username string, password string) (*model.User, string)
	Register(user *model.User) error
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
		user.Password = ""
		return user, "success"
	}
	return nil, "check password failed"
}
func (s *userService) Register(user *model.User) error {
	if u, _ := model.GetUserByUsername(user.Username); u.ID > 0 {
		return &common.CommonError{Message: "用户已存在"}
	}
	user.Salt = util.UUID()
	user.Password = buildPassword(user.Password, user.Salt)
	model.DB.Create(user)
	return nil
}
func checkPassword(user *model.User, password string) bool {
	return buildPassword(password, user.Salt) == user.Password
}

func buildPassword(password string, salt string) string {
	return util.Md5String(password + salt)
}
