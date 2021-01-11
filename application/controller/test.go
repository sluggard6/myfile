package controller

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type TestController struct {
}

func (c *TestController) Get() string {
	// ctx.getu
	return "path not found."
}

func (c *TestController) GetPing(ctx iris.Context) string {
	return "pong"
}

func (c *TestController) GetHello(ctx iris.Context) HttpResult {
	return Success(fmt.Sprintf("Hello %s!", ctx.URLParam("name")))
}
