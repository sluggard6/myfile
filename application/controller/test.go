package controller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type TestController struct {
}

func (c *TestController) Get() string {
	// ctx.getu
	return "path not found."
}

func (c *TestController) GetPing(ctx iris.Context) {
	ctx.WriteString("pong")
	// return "pong"
}

func (m *TestController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // 以及您已经知道的任何标准API调用

	// 1-> Method
	// 2-> Path
	// 3-> 控制器的函数名称将被解析为处理程序
	// 4-> 应该在MyCustomHandler之前运行的任何处理程序
	b.Handle("OPTIONS", "/{any:path}", "", Cors)
}

func (c *TestController) GetHello(ctx iris.Context) *HttpResult {
	return Success(fmt.Sprintf("Hello %s!", ctx.URLParam("name")))
}
