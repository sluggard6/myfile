package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

//FolderController 目录控制器
type FolderController struct {
	folderService service.FolderService
}

//NewFolderController 创建
func NewFolderController() *FolderController {
	return &FolderController{service.NewFolderService()}
}

//Children 目录内的所有内容
type Children struct {
	Folders *[]model.Folder `json:"folders"`
	Files   *[]model.File   `json:"files"`
}

//GetBy 获取目录内搜有内容
func (c *FolderController) GetBy(folderID uint, ctx iris.Context) HttpResult {
	folder := &model.Folder{}
	model.GetById(folder, folderID)
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	if hasRole, _ := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole {
		return FailedForbidden(ctx)
	}
	var folders *[]model.Folder
	var files *[]model.File
	var err error
	folders, err = c.folderService.GetChildrenFolder(folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	files, err = c.folderService.GetChildrenFile(folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	return Success(&Children{folders, files})
}

//PostBy 创建子目录
func (c *FolderController) PostBy(folderID uint, ctx iris.Context) HttpResult {
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	folder, err := service.GetByID(&model.Folder{}, folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	if hasRole, role := user.(*model.User).HasLibrary(folder.(*model.Folder).LibraryID); !hasRole || role == model.Read {
		return FailedForbidden(ctx)
	}
	name := ctx.URLParamDefault("name", "")
	if name == "" {
		return FailedCode(PARAM_ERROR)
	}
	child, err := c.folderService.CreateChild(folder.(*model.Folder), name)
	if err != nil {
		return FailedMessage(err.Error())
	}
	return Success(child)
}
