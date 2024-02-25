package db

import (
	"fmt"

	"toko1/helper"
	"toko1/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(dbUsername, dbPassword, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	//MEMBUAT TABLE DATABASE AUTOMATIS
	db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Product{}, &entity.Cart{}, &entity.Transaction{})

	return db, helper.ReturnIfError(err)
}
