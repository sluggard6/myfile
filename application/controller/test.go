package controller

import "github.com/kataras/iris/v12"

type TestController struct {
}

func (c *TestController) Get() string {
	// ctx.getu
	return "path not found."
}

func (c *TestController) GetPing(ctx iris.Context) string {
	return "pong"
}
