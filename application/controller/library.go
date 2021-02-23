package controller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type LibraryController struct {
	libraryService service.LibraryService
}

func NewLibraryController() *LibraryController {
	return &LibraryController{service.NewLibraryService()}
}

func (c *LibraryController) Get(ctx iris.Context) *HttpResult {
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

func (c *LibraryController) Put(ctx iris.Context) *HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	name := ctx.URLParam("name")
	if user.HasLibraryName(name) {
		return FailedMessage(fmt.Sprintf("资料库'%s'已存在", name))
	}
	library, err := c.libraryService.CreateLibrary(user.ID, name)
	if err != nil {
		return FailedMessage(err.Error())
	}
	user.Librarys = append(user.Librarys, *library)
	return Success(library)
}

func (c *LibraryController) Post(ctx iris.Context) *HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	library := &model.Library{}
	ctx.ReadJSON(library)
	// name := ctx.URLParam("name")
	if user.HasLibraryName(library.Name) {
		return FailedMessage(fmt.Sprintf("资料库'%s'已存在", library.Name))
	}
	err := c.libraryService.UpdateLibrary(library)
	if err != nil {
		return FailedMessage(err.Error())
	}
	return Success(library)
}

func (c *LibraryController) GetCheck(ctx iris.Context) *HttpResult {
	name := ctx.URLParam("name")
	user := sessions.Get(ctx).Get("user").(*model.User)
	if user.HasLibraryName(name) {
		return SuccessMessage(fmt.Sprintf("资料库'%s'已存在", name), false)
	} else {
		return Success(true)
	}
}

func (c *LibraryController) DeleteBy(id uint, ctx iris.Context) *HttpResult {
	user := sessions.Get(ctx).Get("user").(*model.User)
	if b, _ := user.HasLibrary(id); !b {
		return FailedForbidden(ctx)
	}
	err := c.libraryService.DeleteLibrary(id)
	if err != nil {
		return FailedMessage(err.Error())
	}
	// 移除user中已被删除的library
	for i, library := range user.Librarys {
		if library.ID == id {
			user.Librarys = append(user.Librarys[:i], user.Librarys[i+1:]...)
			break
		}
	}
	logrus.Debug(user.Librarys)
	return Success(id)
}

func (c *LibraryController) PutShare() *HttpResult {
	return nil
}
