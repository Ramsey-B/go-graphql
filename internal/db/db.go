package db

import (
	"go-graphql/graph/model"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	err := connect()

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	db.Exec("CREATE DATABASE IF NOT EXISTS test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})

	return db, nil
}

func connect() error {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True" // move this to a config
	db, err = gorm.Open("mysql", dataSourceName)
	return err
}

func setup() {
	db.LogMode(true)

	db.Exec("CREATE DATABASE IF NOT EXISTS test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
}
