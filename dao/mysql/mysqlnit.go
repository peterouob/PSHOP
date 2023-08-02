package mysql

import (
	"PSHOP/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
}
