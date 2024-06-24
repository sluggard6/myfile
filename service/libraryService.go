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
	ShareLibraryOne(libraryID uint, userID uint, role model.LibraryRole) error
	ShareLibrarys(libraryID uint, userIds []uint, role model.LibraryRole) error
	RemoveShareLibrary(shareLibraryID uint, userID uint) error
	// ShareLibrary(shareLibrary *model.ShareLibrary) error
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
	folder.ParentID = folder.ID
	model.UpdateById(folder)
	library.RootFolder = *folder
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

func (s *librarySer) ShareLibrarys(libraryID uint, userIds []uint, role model.LibraryRole) error {
	if len(userIds) == 0 {
		return errors.New("userIds is nil")
	}
	var librarys []model.ShareLibrary
	for _, userID := range userIds {
		// libraryId[i] = model.ShareLibrary{}
		librarys = append(librarys, model.ShareLibrary{UserID: userID, LibraryID: libraryID, Role: role})

	}
	_, err := model.Create(&librarys)
	return err
}

func (s *librarySer) ShareLibraryOne(libraryID uint, userID uint, role model.LibraryRole) error {
	if user, _ := userService.GetById(userID); user.ID <= 0 {
		return fmt.Errorf("can't find user %d", libraryID)
	}
	if library, _ := model.GetById(&model.Library{}, libraryID); library.(*model.Library).ID <= 0 {
		return fmt.Errorf("can't find library %d", libraryID)
	} else {
		if library.(*model.Library).UserID == userID {
			return fmt.Errorf("can't share library to onwer")
		}
	}
	shareLibrary := &model.ShareLibrary{}
	var err error
	if shareLibrary, err = shareLibrary.GetLibraryByUserAndLibrary(userID, libraryID); shareLibrary.ID > 0 {
		if shareLibrary.Role != role {
			shareLibrary.Role = role
			model.UpdateById(shareLibrary)
		}
		// return fmt.Errorf("共享资料库%s(id:%d)已存在", shareLibrary.Library.Name, libraryID)
	} else {
		_, err = model.Create(&model.ShareLibrary{UserID: userID, LibraryID: libraryID, Role: role})
	}
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

func (s *librarySer) RemoveShareLibrary(shareLibraryID uint, userID uint) error {
	shareLibrary := &model.ShareLibrary{ModelHard: model.ModelHard{ID: shareLibraryID}}
	return model.DB().Where("user_id=?", userID).Delete(shareLibrary).Error
}

// func (s *librarySer) ShareLibrary(shareLibrary *model.ShareLibrary) error {
// 	result := db.Where(&ShareLibrary{UserID: l.UserID, LibraryID: l.LibraryID}).First(&shareLibrary)
// 	if err = result.Error; err != nil {
// 		print(err.Error())
// 		return
// 	}
// 	if result.RowsAffected == 0 {
// 		Create(l)
// 		shareLibrary = l
// 	} else {
// 		if shareLibrary.Role != l.Role {
// 			shareLibrary.Role = l.Role
// 			UpdateById(shareLibrary)
// 		}
// 	}
// 	return
// }
