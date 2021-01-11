package controller

type Data interface {
}

type HttpResult struct {
	Code    MessageCode
	Message string
	Data    Data
}

type MessageCode int

const (
	SUCCESS     MessageCode = 0
	FAILED      MessageCode = 1
	PARAM_ERROR MessageCode = 2
)

var failedMessage map[MessageCode]string = map[MessageCode]string{
	SUCCESS:     "success",
	FAILED:      "failed",
	PARAM_ERROR: "param error",
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
