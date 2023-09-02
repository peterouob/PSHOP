package mysql

import (
	user2 "PSHOP/model/user"
	"PSHOP/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var err error

func MysqlInit() {
	dsn := utils.Config.GetString("mysql.dsn")
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		err = fmt.Errorf("database connection have problem: %v", err)
		fmt.Println(err)
	}
	fmt.Println("Mysql connection ...")
	if err := Db.AutoMigrate(
		&user2.UserModel{},
		&user2.Block{},
		&user2.Goods{},
		&user2.GoodInfo{},
		&user2.Comment{},
		&user2.Replay{},
	); err != nil {
		log.Printf("error : %v", err)
	}
	fmt.Println("successfully migrated")
}