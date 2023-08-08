package main

import (
	user2 "PSHOP/model/user"
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
		&user2.UserModel{},
		&user2.Block{},
		&user2.Block{},
	); err != nil {
		log.Printf("error : %v", err)
	}
	fmt.Println("successfully migrated")
}
