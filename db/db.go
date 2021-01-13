package db

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(sqlite.Open("myfile.db"), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	SqlDB, err := db.DB()
	if err != nil {
		log.Error(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	SqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	SqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	SqlDB.SetConnMaxLifetime(time.Hour)

}
