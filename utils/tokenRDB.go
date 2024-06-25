package utils

import (
	"PSHOP/model/database/redis"
	"context"
	"time"
)

func SaveTokenAuth(userid string, td *Token) error {
	at := time.Unix(td.AtExp, 0)
	rt := time.Unix(td.ReExp, 0)
	now := time.Now()
	errAccess := redisdao.Rdb.Set(context.Background(), td.AccessUUid, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefersh := redisdao.Rdb.Set(context.Background(), td.RefreshUUid, td.RefreshToken, rt.Sub(now)).Err()
	if errRefersh != nil {
		return errRefersh
	}
	return nil
}

func DeleteTokenAuth(uid string) (int64, error) {
	deleted, err := redisdao.Rdb.Del(context.Background(), uid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
