package controller

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
)

type FileController struct {
	fileService service.FileService
}

//NewFileController 创建
func NewFileController() *FileController {
	return &FileController{}
}

func (c *FileController) PostUpload(ctx iris.Context) HttpResult {
	file, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return FailedMessage(err.Error())
	}
	folder := &model.Folder{}
	folderID, err := strconv.Atoi(ctx.URLParam("folderID"))
	if err != nil {
		return FailedCode(PARAM_ERROR)
	}
	model.GetById(folder, uint(folderID))
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	if hasRole, role := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole || role == model.Read {
		return FailedForbidden(ctx)
	}
	dbFile, err := c.fileService.SaveFile(file, fileHeader.Filename, folder)
	if err != nil {
		return FailedMessage(err.Error())
	}
	return Success(dbFile)
}
