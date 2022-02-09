package controller

import (
	"fmt"
	"strconv"

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

//GetBy 获取目录内所有内容
func (c *FolderController) GetBy(folderID uint, ctx iris.Context) *HttpResult {
	folder := &model.Folder{}
	model.GetById(folder, folderID)
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	var libraryName string
	var hasRole bool
	if hasRole, _, libraryName = user.(*model.User).HasLibrary(folder.LibraryID); !hasRole {
		return FailedForbidden(ctx)
	}
	var folders *[]model.Folder
	var files *[]model.File
	var path []model.PathItem
	var err error
	folders, err = c.folderService.GetChildrenFolder(folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	files, err = c.folderService.GetChildrenFile(folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	children := &Children{folders, files}
	path, err = folder.GetPath()
	if err != nil {
		return FailedMessage(err.Error())
	}
	ret := struct {
		Folder      *model.Folder    `json:"folder"`
		Children    *Children        `json:"children"`
		Path        []model.PathItem `json:"path"`
		LibraryName string           `json:"libraryName"`
	}{
		folder,
		children,
		path,
		libraryName,
	}
	return Success(ret)
}

//PutBy 创建子目录
func (c *FolderController) PutBy(folderID uint, ctx iris.Context) *HttpResult {
	sess := sessions.Get(ctx)
	user := sess.Get("user").(*model.User)
	folder, err := service.GetByID(&model.Folder{}, folderID)
	if err != nil {
		return FailedMessage(err.Error())
	}
	if hasRole, role, _ := user.HasLibrary(folder.(*model.Folder).LibraryID); !hasRole || role == model.Read {
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

//更新文件夹名称
func (c *FolderController) PostBy(parentId uint, ctx iris.Context) *HttpResult {
	editFolderForm := struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}{}
	user := sessions.Get(ctx).Get("user").(*model.User)
	if err := ctx.ReadJSON(&editFolderForm); err != nil {
		return FailedCodeMessage(PARAM_ERROR, err.Error())
	}
	_folder, err := service.GetByID(&model.Folder{}, editFolderForm.Id)
	if err != nil {
		return FailedMessage(err.Error())
	}
	folder := _folder.(*model.Folder)
	if hasRole, role, _ := user.HasLibrary(folder.LibraryID); !hasRole || role == model.Read {
		return FailedForbidden(ctx)
	}
	folder.Name = editFolderForm.Name
	if err = c.folderService.UpdateFolder(folder); err != nil {
		return FailedMessage(err.Error())
	}
	return Success(folder)
}

// GetCheck 检测资料库是否重名
func (c *FolderController) GetCheck(ctx iris.Context) *HttpResult {
	name := ctx.URLParam("name")
	parentID, err := strconv.Atoi(ctx.URLParam("id"))
	if err != nil {
		return FailedMessage(err.Error())
	}
	folders, err := c.folderService.GetChildrenFolder(uint(parentID))
	if err != nil {
		return FailedMessage(err.Error())
	}
	for _, folder := range *folders {
		if name == folder.Name {
			return SuccessMessage(fmt.Sprintf("文件夹'%s'已存在", name), false)
		}
	}
	return Success(true)
}
