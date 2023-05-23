package main

import (
	"fmt"
	config2 "gin-api-test/config"
	"gin-api-test/db/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {

	ConnectDB()
}
func main() {
	DB.AutoMigrate(&migrations.User{})
	fmt.Println("? Migration complete")
}

var DB *gorm.DB

func ConnectDB() {
	var err error
	config := config2.MyEnv
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	DB.AutoMigrate(&migrations.User{})

	fmt.Println("?Connect database Successfully")
}
