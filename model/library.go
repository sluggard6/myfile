package model

type LibraryRole uint

const (
	Read  LibraryRole = 1
	Write LibraryRole = 2
)

type Library struct {
	Model
	Name       string `json:"name"`
	UserID     uint   `json:"-"`
	Owner      User   `gorm:"foreignKey:UserID" json:"owner"`
	RootFolder Folder `gorm:"foreignKey:LibraryID" json:"rootFolder"`
}

type ShareLibrary struct {
	ModelHard
	UserID    uint        `gorm:"uniqueIndex:un_user_library" json:"userId"`
	LibraryID uint        `gorm:"uniqueIndex:un_user_library" json:"libraryId"`
	Library   Library     `gorm:"foreignKey:ID;references:LibraryID" json:"library"`
	Role      LibraryRole `json:"role"`
}

func (l *Library) GetLibraryMine(userId uint) (librarys []Library, err error) {
	err = db.Where("user_id=?", userId).Preload("RootFolder", "parent_id=id").Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryShare(userId uint) (librarys []ShareLibrary, err error) {
	err = db.Where("user_id=?", userId).Preload("Library").Preload("Library.Owner").Preload("Library.RootFolder").Find(&librarys).Error
	// err = db.Where("user_id=?", userId).Joins("Library").Joins("Library.RootFolder").Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryByUserAndLibrary(userId uint, libraryId uint) (shareLibrary *ShareLibrary, err error) {
	err = db.Where(&ShareLibrary{UserID: userId, LibraryID: libraryId}).First(&shareLibrary).Error
	return
}
