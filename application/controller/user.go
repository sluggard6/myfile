package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController() *UserController {
	return &UserController{service.NewUserService()}
}

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginInfo struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

func (c *UserController) PostLogin(ctx iris.Context) HttpResult {
	loginForm := &LoginForm{}
	if err := ctx.ReadJSON(loginForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return FailedCode(PARAM_ERROR)
	}
	if user, err := c.userService.Login(loginForm.Username, loginForm.Password); user == nil {
		return FailedCodeMessage(LOGIN_FAILED, err.Error())
	} else {
		session := sessions.Get(ctx)
		session.Set("authenticated", true)
		session.Set("user", user)
		return Success(&LoginInfo{User: user, Token: session.ID()})
	}
}

func (c *UserController) GetInfo(ctx iris.Context) HttpResult {
	session := sessions.Get(ctx)
	return c.GetBy(session.Get("user").(*model.User).ID)
}

func (c *UserController) GetBy(id uint) HttpResult {
	if user, err := c.userService.GetById(id); err != nil {
		return Failed()
	} else {
		user.Password = ""
		return Success(user)
	}
}

func (c *UserController) PostRegister(ctx iris.Context) HttpResult {
	user := &model.User{}
	if err := ctx.ReadJSON(user); err != nil {
		return FailedCode(PARAM_ERROR)
	}
	if err := validate.Struct(user); err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return FailedCodeMessage(PARAM_ERROR, err.Error())
		}

		// from here you can create your own error messages in whatever language you wish
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	if err := c.userService.Register(user); err != nil {
		return FailedCodeMessage(LOGIN_FAILED, err.Error())
	}
	return Success(user)
}

func (c *UserController) GetTest(ctx iris.Context) HttpResult {
	return Success("test success")
}
