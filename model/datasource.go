package model

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var cfg config.Database

func Init() error {
	cfg = config.GetConfig().Database
	log.Info(fmt.Sprintf("%s:%s@%s", cfg.Username, cfg.Password, cfg.Url))
	dsn := fmt.Sprintf("%s:%s@%s", cfg.Username, cfg.Password, cfg.Url)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	DB = db
	initTable()
	return nil
}
func initTable() error {
	// log.Debug(reflect.TypeOf(User{}))
	DB.AutoMigrate(&User{})
	return nil
}
