package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root:831374@tcp(127.0.0.1:3306)/messenger?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = DB.AutoMigrate(&Chat{}, &ChatMember{}, &Message{})
	if err != nil {
		log.Printf("failed to migrate: %v", err)
		return
	}
	fmt.Println("Migration done ..")
}
