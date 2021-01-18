package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type Data interface {
}

type HttpResult struct {
	Code    MessageCode
	Message string
	Data    Data
}

type MessageCode int

const (
	SUCCESS      MessageCode = 0
	FAILED       MessageCode = 1
	PARAM_ERROR  MessageCode = 2
	LOGIN_FAILED MessageCode = 101
)

var validate *validator.Validate = validator.New()

var failedMessage map[MessageCode]string = map[MessageCode]string{
	SUCCESS:      "success",
	FAILED:       "failed",
	PARAM_ERROR:  "param error",
	LOGIN_FAILED: "login_failed",
}

func Success(data Data) HttpResult {
	return HttpResult{0, "success", data}
}

func Failed() HttpResult {
	return HttpResult{-1, "failed", nil}
}

func FailedCode(code MessageCode) HttpResult {
	return HttpResult{code, failedMessage[code], nil}
}

func FailedCodeMessage(code MessageCode, message string) HttpResult {
	return HttpResult{Code: code, Message: message, Data: nil}
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
