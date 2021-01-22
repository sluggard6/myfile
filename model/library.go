package model

type LibraryRole uint

const (
	Read  LibraryRole = 1
	Write LibraryRole = 2
)

type Library struct {
	Model
	Name   string `json:"name"`
	UserId uint   `json:"-"`
	Owner  User   `gorm:"foreignKey:UserId"`
}

type ShareLibrary struct {
	Model
	UserId    uint
	LibraryId uint
	Library   Library `gorm:"foreignKey:ID;references:LibraryId"`
	Role      LibraryRole
}

func (l *Library) GetLibraryMine(userId uint) (librarys []Library, err error) {
	err = db.Where("user_id=?", userId).Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryShare(userId uint) (librarys []ShareLibrary, err error) {
	err = db.Where("user_id=?", userId).Find(&librarys).Error
	return
}

func (l *ShareLibrary) GetLibraryByUserAndLibrary(userId uint, libraryId uint) (shareLibrary *ShareLibrary, err error) {
	err = db.Where(&ShareLibrary{UserId: userId, LibraryId: libraryId}).First(shareLibrary).Error
	return
}
