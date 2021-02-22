package service

import (
	"errors"
	"fmt"

	"github.com/sluggard/myfile/model"
)

// LibraryService 资料库服务
type LibraryService interface {
	CreateLibrary(userID uint, name string) (*model.Library, error)
	GetLibraryMine(userID uint) ([]model.Library, error)
	GetLibraryShare(userID uint) ([]model.ShareLibrary, error)
	UpdateLibrary(library *model.Library) error
	DeleteLibrary(id uint) error
}

var libraryService = &librarySer{}

// NewLibraryService 创建LibraryService实现
func NewLibraryService() LibraryService {
	return libraryService
}

type librarySer struct {
}

func (s *librarySer) CreateLibrary(userID uint, name string) (*model.Library, error) {
	if err := checkLibraryName(name); err != nil {
		return nil, err
	}
	_, err := userService.GetById(userID)
	// _, err := model.GetById(&model.User{}, userId)
	if err != nil {
		return nil, err
	}
	library := &model.Library{Name: name, UserID: userID}
	if _, err := model.Create(library); err != nil {
		return nil, err
	}
	folder := &model.Folder{Name: "/", LibraryID: library.ID}
	if _, err := model.Create(folder); err != nil {
		return nil, err
	}
	return library, nil
}
func (s *librarySer) GetLibraryMine(userID uint) ([]model.Library, error) {
	library := &model.Library{}
	return library.GetLibraryMine(userID)
}

func (s *librarySer) GetLibraryShare(userID uint) ([]model.ShareLibrary, error) {
	share := &model.ShareLibrary{}
	return share.GetLibraryShare(userID)
}

func (s *librarySer) ShareLibrary(libraryID uint, userIds []uint) error {
	if len(userIds) == 0 {
		return errors.New("userIds is nil")
	}
	var librarys []model.ShareLibrary
	for _, userID := range userIds {
		// libraryId[i] = model.ShareLibrary{}
		librarys = append(librarys, model.ShareLibrary{UserID: userID, LibraryID: libraryID, Role: model.Write})

	}
	_, err := model.Create(&librarys)
	return err
}

func (s *librarySer) ShareLibraryOne(libraryID uint, userID uint) error {
	if user, _ := userService.GetById(userID); user.ID <= 0 {
		return fmt.Errorf("can't find user %d", libraryID)
	}
	if library, _ := model.GetById(&model.Library{}, libraryID); library.(model.Library).ID <= 0 {
		return fmt.Errorf("can't find library %d", libraryID)
	}
	shareLibrary := &model.ShareLibrary{}
	if shareLibrary, _ = shareLibrary.GetLibraryByUserAndLibrary(userID, libraryID); shareLibrary.ID > 0 {
		return fmt.Errorf("共享资料库%s(id:%d)已存在", shareLibrary.Library.Name, libraryID)
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
