package db

import (
	"fmt"
	"go-graphql/config"
	"go-graphql/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Configuration) (*gorm.DB, error) {
	db, err := connect(cfg)

	if err != nil {
		return nil, err
	}

	setup(cfg, db)

	return db, nil
}

func connect(cfg *config.Configuration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.DBHost, cfg.Database.DBUser, cfg.Database.DBPassword, cfg.Database.DBName, cfg.Database.DBPort)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// this will likely be replaced with a tool like sqitch
func setup(cfg *config.Configuration, db *gorm.DB) {
	db.Exec(fmt.Sprintf("CREATE '%s'", cfg.Database.DBName))

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
}
