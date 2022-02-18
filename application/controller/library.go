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
	if lTpye == "mine" {
		if librarys, err := c.libraryService.GetLibraryMine(user.ID); err != nil {
			return FailedMessage(err.Error())
		} else {
			return Success(librarys)
		}
	} else if lTpye == "share" {
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
	editLibraryForm := struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}{}
	if err := ctx.ReadJSON(&editLibraryForm); err != nil {
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	user := sessions.Get(ctx).Get("user").(*model.User)
	if !user.OwnLibrary(editLibraryForm.Id) {
		return FailedForbidden(ctx)
	}
	_library, err := service.GetByID(&model.Library{}, editLibraryForm.Id)
	library := _library.(*model.Library)
	if user.HasLibraryName(editLibraryForm.Name) {
		return FailedMessage(fmt.Sprintf("资料库'%s'已存在", editLibraryForm.Name))
	}
	if err != nil {
		return FailedMessage(err.Error())
	}
	library.Name = editLibraryForm.Name
	err = c.libraryService.UpdateLibrary(library)
	if err != nil {
		return FailedMessage(err.Error())
	}
	// 更新user中的library
	for i, _library := range user.Librarys {
		if _library.ID == editLibraryForm.Id {
			user.Librarys[i] = *library
			break
		}
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
	if b, _, _ := user.HasLibrary(id); !b {
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

func (c *LibraryController) PutShare(ctx iris.Context) *HttpResult {
	shareLibraryForm := struct {
		LibraryId  uint   `json:"id"`
		Role       uint   `json:"role"`
		ShareUsers []uint `json:"users"`
	}{}
	if err := ctx.ReadJSON(&shareLibraryForm); err != nil {
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	user := sessions.Get(ctx).Get("user").(*model.User)

	if user.OwnLibrary(shareLibraryForm.LibraryId) {
		return FailedForbidden(ctx)
	}
	for userId := range shareLibraryForm.ShareUsers {
		shareLibrary := model.ShareLibrary{
			UserID:    uint(userId),
			LibraryID: shareLibraryForm.LibraryId,
			Role:      model.LibraryRole(shareLibraryForm.Role),
		}
		model.Create(shareLibrary)

	}
	print(user)
	print(shareLibraryForm)
	return nil
}
