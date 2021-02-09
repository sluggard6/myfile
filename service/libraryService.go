package service

import (
	"fmt"

	"github.com/sluggard/myfile/common"
	"github.com/sluggard/myfile/model"
)

type LibraryService interface {
	CreateLibrary(userId uint, name string) (*model.Library, error)
	GetLibraryMine(userId uint) ([]model.Library, error)
	GetLibraryShare(userId uint) ([]model.ShareLibrary, error)
	UpdateLibrary(library *model.Library) error
	DeleteLibrary(id uint) error
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
	_, err := userService.GetById(userId)
	// _, err := model.GetById(&model.User{}, userId)
	if err != nil {
		return nil, err
	}
	library := &model.Library{Name: name, UserID: userId}
	if _, err := model.Create(library); err != nil {
		return nil, err
	}
	folder := &model.Folder{Name: "/", LibraryID: library.ID}
	if _, err := model.Create(folder); err != nil {
		return nil, err
	}
	return library, nil
}
func (s *librarySer) GetLibraryMine(userId uint) ([]model.Library, error) {
	library := &model.Library{}
	return library.GetLibraryMine(userId)
}

func (s *librarySer) GetLibraryShare(userId uint) ([]model.ShareLibrary, error) {
	share := &model.ShareLibrary{}
	return share.GetLibraryShare(userId)
}

func (s *librarySer) ShareLibrary(libraryId uint, userIds []uint) error {
	if len(userIds) == 0 {
		return common.CommonError{"userIds is nil"}
	}
	var librarys []model.ShareLibrary
	for _, userId := range userIds {
		// libraryId[i] = model.ShareLibrary{}
		librarys = append(librarys, model.ShareLibrary{UserID: userId, LibraryID: libraryId, Role: model.Write})

	}
	_, err := model.Create(&librarys)
	return err
}

func (s *librarySer) ShareLibraryOne(libraryID uint, userID uint) error {
	if user, _ := userService.GetById(userID); user.ID <= 0 {
		return common.CommonError{fmt.Sprintf("can't find user %d", libraryID)}
	}
	if library, _ := model.GetById(&model.Library{}, libraryID); library.(model.Library).ID <= 0 {
		return common.CommonError{fmt.Sprintf("can't find library %d", libraryID)}
	}
	shareLibrary := &model.ShareLibrary{}
	if shareLibrary, _ = shareLibrary.GetLibraryByUserAndLibrary(userID, libraryID); shareLibrary.ID > 0 {
		return common.CommonError{Message: fmt.Sprintf("共享资料库%s(id:%d)已存在", shareLibrary.Library.Name, libraryID)}
	}
	_, err := model.Create(&model.ShareLibrary{UserID: userID, LibraryID: libraryID})
	return err
}

func (s *librarySer) DeleteLibrary(id uint) error {
	library := &model.Library{Model: model.Model{ID: id}}
	folderService.DeleteByLibrary(id)
	model.Delete(library)
	return nil
}

func (s *librarySer) UpdateLibrary(library *model.Library) error {
	return model.DB().Save(library).Error
}
