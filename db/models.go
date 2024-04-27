package db

import (
	"time"
)

type Message struct {
	ID       int `gorm:"primary_key"`
	SenderID int `gorm:"foreign_key"`
	ChatID   int `gorm:"foreign_key"`
	Content  string
}

type Gender int8

const (
	Male Gender = iota
	Female
)

type User struct {
	ID          string `gorm:"primary_key"`
	Username    string
	FirstName   string
	LastName    string
	Password    string
	Gender      Gender
	DateOfBirth time.Time
	CreatedTime time.Time
}
