package main

import (
	"PSHOP/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/pshop?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))
	var u model.UserModel
	db.AutoMigrate(&u)
}
