package controller

import (
	"unsafe"

	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/service"
)

type UserController struct {
	userService service.UserService
}

// ABACD
func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}

type LoginForm struct {
	Username string `json:username`
	Password string `json:password`
}

func (c *UserController) PostLogin(ctx iris.Context) HttpResult {
	loginForm := &LoginForm{}
	if err := ctx.ReadJSON(loginForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return FailedCode(PARAM_ERROR)
	}
	log.Debug(loginForm)
	log.Debugf("loginForm.username is %s,loginForm.password is %s", loginForm.Username, loginForm.Password)
	if user, message := c.userService.Login(loginForm.Username, loginForm.Password); user == nil {
		return FailedCodeMessage(LOGIN_FAILED, message)
	} else {
		return Success(user)
	}
}

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
