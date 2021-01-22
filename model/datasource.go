package model

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var cfg config.Database

func Init() error {
	cfg = config.GetConfig().Database
	log.Debug(fmt.Sprintf("%s:%s@%s", cfg.Username, cfg.Password, cfg.Url))
	dsn := fmt.Sprintf("%s:%s@%s", cfg.Username, cfg.Password, cfg.Url)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	// db, err := gorm.Open(sqlite.Open("myfile.db"), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Error(err)
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(2)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Hour)

	// DB = db
	initTable()
	return nil
}
func initTable() error {
	// log.Debug(reflect.TypeOf(User{}))
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Library{})
	db.AutoMigrate(&Folder{})
	db.AutoMigrate(&ShareLibrary{})
	return nil
}
