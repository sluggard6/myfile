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
	Model
	UserID    uint    `json:"userId"`
	LibraryID uint    `json:"libraryId"`
	Library   Library `gorm:"foreignKey:ID;references:LibraryID"`
	Role      LibraryRole
}

func (l *Library) GetLibraryMine(userId uint) (librarys []Library, err error) {
	err = db.Where("user_id=?", userId).Preload("RootFolder", "parent_id=?", 0).Find(&librarys).Error
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
