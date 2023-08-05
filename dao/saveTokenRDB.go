package dao

import (
	"PSHOP/dao/redis"
	"PSHOP/utils"
	"context"
	"time"
)

func SaveTokenAuth(userid string, td *utils.Token) error {
	at := time.Unix(td.AtExp, 0)
	rt := time.Unix(td.ReExp, 0)
	now := time.Now()
	errAccess := redis.Rdb.Set(context.Background(), td.AccessUUid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefersh := redis.Rdb.Set(context.Background(), td.RefreshUUid, userid, rt.Sub(now)).Err()
	if errRefersh != nil {
		return errRefersh
	}
	return nil
}
