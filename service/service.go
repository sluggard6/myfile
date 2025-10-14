package service

import "github.com/sluggard/myfile/model"

//GetByID 统一的根据id查询对象方法
func GetByID(modelEntity interface{}, id uint) (interface{}, error) {
	return model.GetById(modelEntity, id)
}

type ServiceGroup struct {
	UserService
	LibraryService
	FolderService
	FileService
	TokenService
}

var ServiceGroupApp = ServiceGroup{
	TokenService: *tokenService,
}
