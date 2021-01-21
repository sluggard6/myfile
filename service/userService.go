package service

import (
	"fmt"
	"strings"

	"github.com/sluggard/myfile/common"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/util"
)

type UserService interface {
	Login(username string, password string) (*model.User, error)
	Register(user *model.User) error
	GetById(id uint) (*model.User, error)
	CreateLibrary(userId uint, name string) (*model.Library, error)
}

var userServiceIntence = &userService{}

func NewUserService() UserService {
	return userServiceIntence
}

type userService struct {
	// repo repositories.MovieRepository
	// userDao dao.UserDao
}

func (s *userService) Login(username string, password string) (user *model.User, err error) {
	if user, err = user.GetUserByUsername(username); err != nil {
		return nil, err
	}
	if checkPassword(user, password) {
		user.Password = ""
		return user, nil
	}
	return nil, &common.CommonError{Message: "check password failed"}
}
func (s *userService) Register(user *model.User) error {
	if user, _ := user.GetUserByUsername(user.Username); user.ID > 0 {
		return &common.CommonError{Message: "用户已存在"}
	}
	user.Salt = util.UUID()
	user.Password = buildPassword(user.Password, user.Salt)
	// model.DB.Create(user)
	model.Create(user)
	s.CreateLibrary(user.ID, "Default Library")
	return nil
}
func (s *userService) GetById(id uint) (user *model.User, err error) {
	return user.GetUserById(id)
	// if user, err = model.DB.Where("id=?", id).Find(user); user.ID > 0 {
	// 	return user, nil
	// } else {
	// 	return nil, err
	// }
}
func (s *userService) CreateLibrary(userId uint, name string) (*model.Library, error) {
	if err := checkLibraryName(name); err != nil {
		return nil, err
	}
	_, err := s.GetById(userId)
	if err != nil {
		return nil, err
	}
	library := &model.Library{Name: name, UserId: userId}
	if _, err := model.Create(library); err != nil {
		return nil, err
	}
	folder := &model.Folder{Name: "/", LibraryId: library.ID}
	if _, err := model.Create(folder); err != nil {
		return nil, err
	}
	role := &model.UserLibraryRole{UserId: userId, LibraryId: library.ID, Role: model.Owner}
	if _, err := model.Create(role); err != nil {
		return nil, err
	}
	return library, nil
}

func (s *userService) reNameLibrary(userId uint, libraryId uint, name string) error {

	return nil
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
