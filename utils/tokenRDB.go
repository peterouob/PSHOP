package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(id string) (string, error) {
	tokenVal := Config.GetString("token.val")
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["user_id"] = id
	claim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tk.SignedString([]byte(tokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return "", err
	}
	return token, nil
}
