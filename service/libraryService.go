package service

import "github.com/sluggard/myfile/model"

func CreateLibrary(userId uint, name string) (*model.Library, error) {
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
