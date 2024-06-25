package utils

import (
	"PSHOP/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
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
type AccessDetails struct {
	AccessUid string
	Userid    uint64
}

var err error

func CreateToken(id uint64) (*Token, error) {
	t := &Token{}
	t.AccessUUid = uuid.NewString()
	t.RefreshUUid = uuid.NewString()
	t.AtExp = time.Now().Add(time.Minute * 15).Unix()
	t.ReExp = time.Now().Add(time.Hour * 24).Unix()
	tokenVal := config.Config.GetString("token.val")
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
	rtokenVal := config.Config.GetString("token.refreshval")
	rclaim := jwt.MapClaims{}
	rclaim["authorized"] = true
	rclaim["refresh_uuid"] = t.RefreshUUid
	rclaim["user_id"] = id
	rclaim["exp"] = time.Now().Add(time.Minute * 2).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.RefreshToken, err = rt.SignedString([]byte(rtokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}
	go refreshTokenRoutine(t.RefreshToken, id, rtokenVal)

	return t, nil
}

// refreshTokenRoutine 定期检查刷新令牌的有效期并自动刷新
func refreshTokenRoutine(refreshToken string, userId uint64, refreshTokenSecret string) {
	for {
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(refreshTokenSecret), nil
		})
		if err != nil {
			fmt.Println("error parsing refresh token:", err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			// 提前 1 分钟刷新
			sleepDuration := time.Duration(exp-time.Now().Unix()-60) * time.Second
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}

			// 刷新令牌
			newToken, err := CreateToken(userId)
			if err != nil {
				fmt.Println("error creating new token:", err)
				return
			}
			refreshToken = newToken.RefreshToken
			fmt.Println("refresh token refreshed successfully")

			// 存储新令牌到 Redis
			saveErr := SaveTokenAuth(strconv.FormatUint(userId, 10), newToken)
			if saveErr != nil {
				fmt.Println("error saving token to redis:", saveErr)
				return
			}
			fmt.Println("new tokens saved to redis successfully")
		} else {
			fmt.Println("invalid refresh token")
			return
		}
	}
}

// ExtractToken 從請求標頭擷取token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		fmt.Println(strArr[1])
		return strArr[1]
	}
	return ""
}

// VerifyToken 檢查token簽名方法
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", tk.Header["alg"])
		}
		return []byte(config.Config.GetString("token.val")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid 驗證token時效性
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata 提取token原數據已進行redis查找
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUid: accessUuid,
			Userid:    userId,
		}, nil
	}
	return nil, err
}
