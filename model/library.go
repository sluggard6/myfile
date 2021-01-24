package model

type LibraryRole uint

const (
	Read  LibraryRole = 1
	Write LibraryRole = 2
)

type Library struct {
	Model
	Name       string `json:"name"`
	UserId     uint   `json:"-"`
	Owner      User   `gorm:"foreignKey:UserId" json:"owner"`
	RootFolder Folder `gorm:"foreignKey:LibraryId" json:"rootFolder"`
}

type ShareLibrary struct {
	Model
	UserID    uint    `json:"userId"`
	LibraryID uint    `json:"libraryId"`
	Library   Library `gorm:"foreignKey:ID;references:LibraryId"`
	Role      LibraryRole
}

func (l *Library) GetLibraryMine(userId uint) (librarys []Library, err error) {
	err = db.Where("user_id=?", userId).Preload("RootFolder").Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryShare(userId uint) (librarys []ShareLibrary, err error) {
	err = db.Where("user_id=?", userId).Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryByUserAndLibrary(userId uint, libraryId uint) (shareLibrary *ShareLibrary, err error) {
	err = db.Where(&ShareLibrary{UserID: userId, LibraryID: libraryId}).First(shareLibrary).Error
	return
}
