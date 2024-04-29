package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *Database

type Database struct {
	db *gorm.DB
}

func NewDatabase() {
	log.Println("Database created")
	var err error
	dsn := "root:831374@tcp(127.0.0.1:3306)/messenger?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = gormDB.AutoMigrate(&Chat{}, &ChatMember{}, &Message{}, &UserTable{})
	if err != nil {
		log.Printf("failed to migrate: %v", err)
		return
	}
	fmt.Println("Migration done ..")
	Mysql.db = gormDB
}
