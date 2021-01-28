package service

import "github.com/sluggard/myfile/model"

func GetById(modelEntity interface{}, id uint) (interface{}, error) {
	return model.GetById(modelEntity, id)
}
