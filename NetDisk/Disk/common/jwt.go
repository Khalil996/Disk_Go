package common

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	NetdiskClaims struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		jwt.RegisteredClaims
	}
)

var secretKey = []byte("netdisk")

// 授权
func GenerateToken(id int64, name string) (string, error) {
	var now = time.Now().Local()
	claims := &NetdiskClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用指定的签名方法和声明创建一个新token
	tokenString, err := token.SignedString(secretKey)          // 创建并返回一个完整的token（jwt）。令牌使用令牌中指定的签名 方法进行签名 。
	return tokenString, err
}

// 鉴权
func AnalyzeToken(token string) (*NetdiskClaims, error) {
	var claims = new(NetdiskClaims)
	// ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	// 第一个值是token ，
	// 第二个值是我们之后需要把解析的数据放入的地方，
	// 第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	tokenClaim, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

	if err != nil || !tokenClaim.Valid {
		return nil, err
	}
	return claims, nil
}
