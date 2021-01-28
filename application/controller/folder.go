package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type FolderController struct {
	folderService service.FolderService
}

func NewFolderController() *FolderController {
	return &FolderController{service.NewFolderService()}
}

type Children struct {
	Folders *[]model.Folder `json:"folders"`
	Files   *[]model.File   `json:"files"`
}

func (c *FolderController) GetBy(folderId uint, ctx iris.Context) HttpResult {
	folder := &model.Folder{}
	model.GetById(folder, folderId)
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	if hasRole, _ := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole {
		return FailedForbidden(ctx)
	}
	var folders *[]model.Folder
	var files *[]model.File
	var err error
	folders, err = c.folderService.GetChildrenFolder(folderId)
	if err != nil {
		FailedMessage(err.Error())
	}
	files, err = c.folderService.GetChildrenFile(folderId)
	if err != nil {
		FailedMessage(err.Error())
	}
	return Success(&Children{folders, files})
}

func (c *FolderController) PostBy(folderId uint, ctx iris.Context) HttpResult {
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	// var folder *model.Folder
	// var err error
	folder, err := service.GetById(&model.Folder{}, folderId)
	if err != nil {
		return FailedMessage(err.Error())
	}
	if hasRole, role := user.(*model.User).HasLibrary(folder.(*model.Folder).LibraryID); !hasRole || role == model.Read {
		return FailedForbidden(ctx)
	}
	if name := ctx.URLParamDefault("name", ""); name == "" {
		return FailedCode(PARAM_ERROR)
	} else {
		child, err := folder.(*model.Folder).CreateChild(name)
		if err != nil {
			return FailedMessage(err.Error())
		}
		return Success(child)
	}
	// if err := ctx.ReadJSON(folder); err != nil {
	// 	return FailedMessage(err.Error())
	// }
	// folder.CreateChild()

	// return Success(folder)
}
