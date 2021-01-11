package controller

import "github.com/kataras/iris/v12"

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
	return Failed()
}
