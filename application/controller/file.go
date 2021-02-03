package controller

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/service"
	"github.com/sluggard/myfile/store"
)

type FileController struct {
	fileService service.FileService
}

//NewFileController 创建
func NewFileController(store store.Store) *FileController {
	return &FileController{service.NewFileService(store)}
}

func (c *FileController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:uint}/{name:string}", "LoadFile")
}

func (c *FileController) LoadFile(ctx iris.Context, id uint, name string) *HttpResult {
	file := &model.File{}
	_, err := service.GetByID(file, id)
	if err != nil {
		return nil
	}
	folder := &model.Folder{}
	_, err = service.GetByID(folder, file.FolderID)
	if err != nil {
		return nil
	}
	sess := sessions.Get(ctx)
	user := sess.Get("user")
	if hasRole, _ := user.(*model.User).HasLibrary(folder.LibraryID); !hasRole {
		ret := FailedForbidden(ctx)
		return &ret
	}
	_, err = service.GetByID(file.Policy, file.PolicyID)
	// src, err := os.Open(file.Policy.Path)
	// if err != nil {
	// 	return nil
	// }
	// defer src.Close()
	// io.Copy(ctx.ResponseWriter(), src)
	ctx.SendFile(file.Policy.Path, file.Name)
	return nil
}

func (c *FileController) PostUpload(ctx iris.Context) HttpResult {
	file, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return FailedMessage(err.Error())
	}
	folder := &model.Folder{}
	folderID, err := strconv.Atoi(ctx.FormValue("folderId"))
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
