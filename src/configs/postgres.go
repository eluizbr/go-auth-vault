package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/eluizbr/go_pagamentos/auth/src/auth/models"
	costumers "github.com/eluizbr/go_pagamentos/auth/src/costumers/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	sslmode  string
}

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsnConfig := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		sslmode:  os.Getenv("DB_SSLMODE"),
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dsnConfig.Host, dsnConfig.Port, dsnConfig.User, dsnConfig.DBName, dsnConfig.Password, dsnConfig.sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(models.User{})
	db.AutoMigrate(costumers.Costumer{})
	db.AutoMigrate(costumers.Address{})
	db.AutoMigrate(costumers.Document{})

	DB = db
}
