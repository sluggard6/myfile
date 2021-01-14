package dao

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqlDb gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("myfile.db"), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Error(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(2)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Hour)

}
