// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"

// 	"github.com/mhghw/fara-message/db"
// )

// type Gender int8

// const (
// 	Male Gender = iota
// 	Female
// )

// // type User struct {
// // 	ID          string
// // 	Username    string
// // 	FirstName   string
// // 	LastName    string
// // 	Password    string
// // 	Gender      Gender
// // 	DateOfBirth time.Time
// // 	CreatedTime time.Time
// // }

// func main() {
// 	// fmt.Println(db.UsersDB)

// 	// var us User
// 	for i := 10001; i <= 10004; i++ {
// 		// var temp string
// 		temp := strconv.Itoa(i)
// 		us, err := db.UsersDB.GetUser(temp)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(reflect.TypeOf(us))
// 	}

// }
