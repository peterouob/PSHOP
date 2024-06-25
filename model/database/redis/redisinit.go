package redisdao

import (
	"PSHOP/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func RedisInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.GetString("redis.addr"),
		Password: "",
		DB:       0,
	})

	ping := Rdb.Ping(context.Background())
	//fmt.Println(ping)
	if ping.String() != "" {
		fmt.Println("Redis connection ...")
	}
}
