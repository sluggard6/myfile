package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/config"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type HttpResult struct {
	Code    MessageCode `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MessageCode int

const (
	SUCCESS      MessageCode = 0
	FAILED       MessageCode = 1
	PARAM_ERROR  MessageCode = 2
	LOGIN_FAILED MessageCode = 101
	// ctx.StatusCode(iris.StatusForbidden)
)

var validate *validator.Validate = validator.New()

var failedMessage map[MessageCode]string = map[MessageCode]string{
	SUCCESS:      "success",
	FAILED:       "failed",
	PARAM_ERROR:  "param error",
	LOGIN_FAILED: "login_failed",
}

func Success(data interface{}) *HttpResult {
	return &HttpResult{SUCCESS, "success", data}
}

func SuccessMessage(message string, data interface{}) *HttpResult {
	return &HttpResult{SUCCESS, message, data}
}

func Failed() *HttpResult {
	return &HttpResult{FAILED, "failed", nil}
}

func FailedMessage(message string) *HttpResult {
	return &HttpResult{FAILED, message, nil}
}

func FailedCode(code MessageCode) *HttpResult {
	return &HttpResult{code, failedMessage[code], nil}
}

func FailedCodeMessage(code MessageCode, message string) *HttpResult {
	return &HttpResult{Code: code, Message: message, Data: nil}
}

func FailedForbidden(ctx iris.Context) *HttpResult {
	ctx.StatusCode(iris.StatusForbidden)
	return &HttpResult{Code: 401, Message: "forbidden"}
}
func Cors(ctx iris.Context) {
	origin := ctx.GetHeader("Origin")
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, X-Token")
		//204
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.Next()
}

func CurrentUser(ctx iris.Context) (user *model.User) {
	if config.GetConfig().Server.AuthType == "token" {
		return ctx.Values().Get("user").(*model.User)
	} else if config.GetConfig().Server.AuthType == "session" {
		return sessions.Get(ctx).Get("user").(*model.User)
	} else {
		return nil
	}
}

type ControllerGroup struct {
	TestController
	UserController
	FolderController
	LibraryController
	FileController
}

func (cg *ControllerGroup) InitTest(party iris.Party) {
	party.Handle("GET", "/ping", testRouter.GetPing)
}

var ControllerGroupApp = new(ControllerGroup)

var (
	testRouter     = ControllerGroupApp.TestController
	userService    = service.ServiceGroupApp.UserService
	tokenService   = service.ServiceGroupApp.TokenService
	folderService  = service.ServiceGroupApp.FolderService
	libraryService = service.ServiceGroupApp.LibraryService
)
