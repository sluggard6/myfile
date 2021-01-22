package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type LibraryController struct {
	libraryService service.LibraryService
}

func NewLibraryController() *LibraryController {
	return &LibraryController{service.NewLibraryService()}
}

func (c *LibraryController) Get(ctx iris.Context) HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	lTpye := ctx.URLParamDefault("type", "mine")
	if "mine" == lTpye {
		if librarys, err := c.libraryService.GetLibraryMine(user.ID); err != nil {
			return FailedMessage(err.Error())
		} else {
			return Success(librarys)
		}
	} else if "share" == lTpye {
		if librarys, err := c.libraryService.GetLibraryShare(user.ID); err != nil {
			return FailedMessage(err.Error())
		} else {
			return Success(librarys)
		}
	} else {
		return FailedCode(PARAM_ERROR)
	}
}
