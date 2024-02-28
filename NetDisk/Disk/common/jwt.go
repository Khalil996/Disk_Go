package common

import (
	"cloud_go/Disk/define"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 授权
func GenerateToken(id int64, name string, second int) (string, error) {
	claims := define.UserClaim{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 鉴权
func AnalyzeToken(token string) (*define.UserClaim, error) {
	claims := new(define.UserClaim)
	userclaim, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !userclaim.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
