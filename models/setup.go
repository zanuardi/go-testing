package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	dsn := "root:zanuardi@tcp(127.0.0.1:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connecting database failed")
	}

	db.AutoMigrate(&User{})

	return db
}
