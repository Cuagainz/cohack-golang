package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

const ConnectionStringFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

var db *gorm.DB

func InitModel() {
	connStr := fmt.Sprintf(ConnectionStringFormat, "root", "test123", "db", 3306, "todo-service")

	var err error
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		logrus.Fatal(err)
	}

	// migration on all the tables
	migrate()
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func migrate() {
	db.AutoMigrate(&Todo{})
}
