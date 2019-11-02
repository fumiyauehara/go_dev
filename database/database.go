package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
}

func connect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "example"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "test"
	ParseTime := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + ParseTime
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}

func (u User) StoreUser() {
	db := connect()
	db.AutoMigrate(&User{})
	db.Create(&u)
	db.Close()
}

func (u User) FindUser(id int) *User{
	db := connect()
	db.First(&u, id)
	db.Close()
	return &u
}
