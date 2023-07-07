package config

import (
	"fmt"
	"golang-jwttoken/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	// sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	sqlInfo := "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
