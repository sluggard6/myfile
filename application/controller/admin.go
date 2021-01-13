package controller

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/model"
)

type AdminController struct {
}

type LoginForm struct {
	username string
	password string
}

func (c *AdminController) PostLogin(ctx iris.Context) HttpResult {
	loginForm := &LoginForm{}
	if err := ctx.ReadJSON(loginForm); err != nil {
		// logging.ErrorLogger.Errorf("login read request json err ", err)
		// ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return FailedCode(PARAM_ERROR)
	}
	log.Debug(loginForm)
	if loginForm.username == "admin" {
		return Success(model.Admin{Name: "admin"})
	}
	return Failed()
}
