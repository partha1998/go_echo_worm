package db

import (
	"fmt"
	"go_echo_gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("successfully conect to the database ")

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // Close
	// defer sqlDB.Close()

	db.AutoMigrate(&models.Product{})

	Database = DbInstance{Db: db}
}
