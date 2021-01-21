package controller

import "github.com/sluggard/myfile/service"

type LibraryController struct {
	libraryService service.LibraryService
}

func NewLibraryController() *LibraryController {
	return &LibraryController{service.NewLibraryService()}
}
