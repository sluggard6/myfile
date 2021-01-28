package model

import (
	"fmt"
	"time"

	"github.com/sluggard/myfile/util"
	"gorm.io/gorm"
)

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// func (this JsonTime) Value() (time.Time, error) {
// 	return time.Time(this), nil
// }

func (this JsonTime) Scan(src interface{}) error {
	this = src.(JsonTime)
	return nil
}

func (this JsonTime) UnmarshalJSON(b []byte) error {
	// var err error
	t, err := time.Parse("2006-01-02 15:04:05", util.ByteArrayToString(b))
	this = JsonTime(t)
	return err
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt JsonTime       `json:"created_at"`
	UpdatedAt JsonTime       `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func Create(model interface{}) (int64, error) {
	result := db.Create(model)
	return result.RowsAffected, result.Error
}

func GetById(model interface{}, id uint) (interface{}, error) {
	// result := db.First(model, id)
	result := db.Where("id = ?", id).First(model)
	return model, result.Error
}

func UpdateById(model interface{}) (int64, error) {
	result := db.Model(model).Updates(model)
	return result.RowsAffected, result.Error
}

func Delete(model interface{}) (int64, error) {
	result := db.Delete(model)
	return result.RowsAffected, result.Error
}
