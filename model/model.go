package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
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
