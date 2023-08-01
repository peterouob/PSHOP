package utils

import "github.com/go-sql-driver/mysql"

var ( // ErrDuplicateEntryCode 命中唯一索引
	ErrDuplicateEntryCode = 1062
)

func MysqlErrcode(err error) int {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return 0
	}
	return int(mysqlErr.Number)
}
