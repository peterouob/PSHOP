package main

import (
	"PSHOP/model/dao/user"
	"fmt"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/pshop?charset=utf8mb4&parseTime=True&loc=Local"
	D, err := gorm.Open(mysql2.Open(dsn))
	if err != nil {

	}
	if err := D.AutoMigrate(
		&user.UserModel{},
		&user.Block{},
		&user.Block{},
	); err != nil {
		log.Printf("error : %v", err)
	}
	fmt.Println("successfully migrated")
}
