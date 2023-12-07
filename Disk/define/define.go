package define

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	Id   int
	Name string
	jwt.RegisteredClaims
}

var JwtKey = "Disk-key"

// token过期时间
var TokenExpire = 3600

// token刷新过期时间
var RefreshTokenExpire = 7200

// 验证码长度
var CodeLength = 6

// 验证码过期时间
var CodeExpire = 300
