package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// d, err := gorm.Open("mysql", "kain:kernelKain999/simplerest?charset=utf8&parseTime=True&loc=Local")
	// if err != nil{
	//     panic(err)
	// }
	// Replace these credentials with your actual MySQL setup
	dsn := "root:BaBaBaBlackSheep283@tcp(127.0.0.1:3306)/bookstore_db?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	log.Println("✅ Successfully connected to MySQL database!")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
