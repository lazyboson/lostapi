package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host         string
	DatabasePort string
	Password     string
	UserName     string
	DatabaseName string
	SSLMode      string
}

func NewConnection(config *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=%s",
		config.Host, config.DatabasePort, config.UserName, config.Password, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	createDBCommand := fmt.Sprintf("CREATE DATABASE %s;", config.DatabaseName)
	db.Exec(createDBCommand)
	return db, nil
}
