package db

import (
	"fmt"
	"go-graphql/config"
	"go-graphql/graph/model"

	"github.com/jinzhu/gorm"
)

func Init(cfg *config.Configuration) (*gorm.DB, error) {
	db, err := connect(cfg)

	if err != nil {
		return nil, err
	}

	setup(db)

	return db, nil
}

func connect(cfg *config.Configuration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.DBHost, cfg.Database.DBUser, cfg.Database.DBPassword, cfg.Database.DBName, cfg.Database.DBPort)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func setup(db *gorm.DB) {
	db.LogMode(true)

	db.Exec("CREATE DATABASE IF NOT EXISTS test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
}
