package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string

	//定義uuid和有效期限->redis儲存時用到
	AccessUUid  string
	RefreshUUid string
	AtExp       int64
	ReExp       int64
}

var err error

func CreateToken(id string) (*Token, error) {
	t := &Token{}
	t.AccessUUid = uuid.NewString()
	t.RefreshUUid = uuid.NewString()
	t.AtExp = time.Now().Add(time.Minute * 15).Unix()
	t.ReExp = time.Now().Add(time.Hour * 24).Unix()
	tokenVal := Config.GetString("token.val")
	// Create token
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["access_uuid"] = t.AccessUUid
	claim["user_id"] = id
	claim["exp"] = t.AtExp
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.AccessToken, err = tk.SignedString([]byte(tokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}
	//Create RefreshToken
	rtokenVal := Config.GetString("token.refreshval")
	rclaim := jwt.MapClaims{}
	rclaim["authorized"] = true
	rclaim["refresh_uuid"] = t.RefreshUUid
	rclaim["user_id"] = id
	rclaim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.RefreshToken, err = rt.SignedString([]byte(rtokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}
	return t, nil
}
