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
	library := &model.Library{Name: name, UserId: userId}
	if _, err := model.Create(library); err != nil {
		return nil, err
	}
	folder := &model.Folder{Name: "/", LibraryId: library.ID}
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

func (s *librarySer) ShareLibraryOne(libraryId uint, userId uint) error {
	if user, _ := userService.GetById(userId); user.ID <= 0 {
		return common.CommonError{fmt.Sprintf("can't find user %s", libraryId)}
	}
	if library, _ := model.GetById(&model.Library{}, libraryId); library.(model.Library).ID <= 0 {
		return common.CommonError{fmt.Sprintf("can't find library %s", libraryId)}
	}
	shareLibrary := &model.ShareLibrary{}
	if shareLibrary, _ = shareLibrary.GetLibraryByUserAndLibrary(userId, libraryId); shareLibrary.ID > 0 {
		return common.CommonError{fmt.Sprintf("共享资料库已存在", libraryId)}
	}
	_, err := model.Create(&model.ShareLibrary{UserID: userId, LibraryID: libraryId})
	return err
}
