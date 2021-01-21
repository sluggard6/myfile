package service

import "github.com/sluggard/myfile/model"

type LibraryService interface {
	CreateLibrary(userId uint, name string) (*model.Library, error)
	getLibraryMine(userId uint) ([]model.Library, error)
	getLibrarShare(userId uint) ([]model.Library, error)
}

var libraryService = &librarySer{}

func NewLibraryService() LibraryService {
	return libraryService
}

type librarySer struct {
}

func (s *librarySer) CreateLibrary(userId uint, name string) (*model.Library, error) {
	if err := checkLibraryName(name); err != nil {
		return nil, err
	}
	_, err := model.GetById(model.User{}, userId)
	if err != nil {
		return nil, err
	}
	library := &model.Library{Name: name, UserId: userId}
	if _, err := model.Create(library); err != nil {
		return nil, err
	}
	folder := &model.Folder{Name: "/", LibraryId: library.ID}
	if _, err := model.Create(folder); err != nil {
		return nil, err
	}
	role := &model.UserLibraryRole{UserId: userId, LibraryId: library.ID, Role: model.Owner}
	if _, err := model.Create(role); err != nil {
		return nil, err
	}
	return library, nil
}
func (s *librarySer) getLibraryMine(userId uint) ([]model.Library, error) {
	library := &model.Library{}
	return library.GetLibraryMine(userId)
}
func (s *librarySer) getLibrarShare(userId uint) ([]model.Library, error) {
	return nil, nil
}
