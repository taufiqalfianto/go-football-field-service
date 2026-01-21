package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	config := Config

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnections)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxLifeTimeConnection) * time.Second)

	return db, nil
}
