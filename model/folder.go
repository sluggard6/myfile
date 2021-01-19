package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Name     string
	PraentId uint
	UserId   uint
}

func (f *Folder) Create() (folder *Folder, err error) {
	if err := DB.Create(folder).Error; err != nil {
		log.Warning("无法插入目录记录, %s", err)
		return nil, err
	}
	return folder, nil
}
