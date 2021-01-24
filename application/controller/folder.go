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
	Folders *[]model.Folder
	Files   *[]model.File
}

func (c *FolderController) GetBy(folderId uint, ctx iris.Context) HttpResult {
	folder := &model.Folder{}
	model.GetById(folder, folderId)
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	if !user.(*model.User).HasLibrary(folder.LibraryId) {
		return FailedForbidden(ctx)
	}
	folders, _ := c.folderService.GetChildrenFolder(folderId)
	files, _ := c.folderService.GetChildrenFile(folderId)
	return Success(Children{folders, files})
}
