package controller

import (
	"unsafe"

	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
)

type AdminController struct {
}

type LoginForm struct {
	username string `json:username`
	password string `json:password`
}

func (c *AdminController) PostLogin(ctx iris.Context) HttpResult {
	var loginForm LoginForm
	if err := ctx.ReadJSON(&loginForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		// logging.ErrorLogger.Errorf("login read request json err ", err)
		// ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return FailedCode(PARAM_ERROR)
	}
	// body, err := ctx.GetBody()
	// if err != nil {
	// 	log.Error(err)
	// }
	// log.Debugf(String(body))
	log.Debugf("loginForm.username is %s,loginForm.password is %s", loginForm.username, loginForm.password)
	return Failed()
}

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
