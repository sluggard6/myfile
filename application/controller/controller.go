package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
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

func Success(data interface{}) HttpResult {
	return HttpResult{SUCCESS, "success", data}
}

func Failed() HttpResult {
	return HttpResult{FAILED, "failed", nil}
}

func FailedMessage(message string) HttpResult {
	return HttpResult{FAILED, message, nil}
}

func FailedCode(code MessageCode) HttpResult {
	return HttpResult{code, failedMessage[code], nil}
}

func FailedCodeMessage(code MessageCode, message string) HttpResult {
	return HttpResult{Code: code, Message: message, Data: nil}
}

func FailedForbidden(ctx iris.Context) HttpResult {
	ctx.StatusCode(iris.StatusForbidden)
	return HttpResult{Code: 401, Message: "forbidden"}
}
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:9528")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, X-Token")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
