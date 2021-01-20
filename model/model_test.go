package model

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// driver "gorm.io/driver/sqlmock"
	"gorm.io/driver/sqLite"
	"gorm.io/gorm"
)

var mock sqlmock.Sqlmock

// TestMain 初始化数据库Mock
func TestMain(m *testing.M) {
	var err error
	var mockDb *sql.DB
	mockDb, mock, err = sqlmock.New()
	fmt.Fprintln(mockDb.Close().Error())
	if err != nil {
		panic("An error was not expected when opening a stub database connection")
	}
	// db, _ = gorm.Open(driver.Open(sqLite.Open("testdb"))
	db, _ = gorm.Open(sqLite.Open("testdb"))
	defer db.Close()
	m.Run()
}
