package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/astaxie/beego"
	"strconv"
)

type EasyToken struct {
	Uid      int64
	Expires  int64
}

var (
	verifyKey  string
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func init() {
	verifyKey = beego.AppConfig.String("jwt_token")
}

func (e EasyToken) GetToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires, //time.Unix(c.ExpiresAt, 0)
		Issuer:    strconv.Itoa(int(e.Uid)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString, err
}

func (e EasyToken) ValidateToken(tokenString string, newExpireAt int64) (bool, error) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})

	if token == nil {
		return false, errors.New(ErrInvalid)
	}
	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		if int64(claims["exp"].(float64)) < newExpireAt {  // 留一个服务端让用户登出的口子
			return false, errors.New(ErrExpired)
		}
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired)
		} else {
			return false, errors.New(ErrOther)
		}
	} else {
		return false, errors.New(ErrOther)
	}
}

// GetUserId from token
func (e EasyToken) GetUserId(tokenString string) (int64) {
	if tokenString == "" {
		return 0
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})

	if token == nil {
		return 0
	}
	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		uid, err:= strconv.Atoi(claims["iss"].(string))
		if err!=nil{
			return 0
		}
		return int64(uid)
	}
	return 0
}