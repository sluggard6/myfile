package model

type User struct {
	Model
	Username      string         `gorm:"unique_index:idx_only_one;commit:'用户名'" validate:"required"`
	Password      string         `gorm:"not null;commit:'用户密码'" validate:"required"`
	Salt          string         `grom:"not null;commit:'用户掩码'" json:"-"`
	Librarys      []Library      `gorm:"foreignKey:UserId"`
	ShareLibrarys []ShareLibrary `gorm:"foreignKey:UserId"`
}

func (u *User) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	result := db.Where("username = ?", username).Preload("Librarys").Preload("ShareLibrarys").First(user)
	return user, result.Error
}

func (u *User) GetUserById(id uint) (*User, error) {
	user := &User{}
	result := db.Where("id = ?", id).First(user)
	return user, result.Error
}

func (u *User) HasLibrary(libraryId uint) bool {
	// db.Joins("Librarys").Joins("ShareLibrarys").Find(u, u.ID)
	for _, library := range u.Librarys {
		if library.ID == libraryId {
			return true
		}
	}
	for _, shareLibrary := range u.ShareLibrarys {
		if shareLibrary.LibraryID == libraryId {
			return true
		}
	}
	return false
}

// func (u *User) CreateUser() (int64, error) {
// 	result := db.Create(u)
// 	return result.RowsAffected, result.Error
// }
