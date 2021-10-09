package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"leek-api/app/models/user"
	"leek-api/pkg/config"
	"time"
)

var jwtSecret = []byte(config.GetString("jwt.secret"))

type Claims struct {
	user.User
	jwt.StandardClaims
}

// GenerateToken 生成令牌
func GenerateToken(user user.User) (string, error) {
	// 过期时间，时间戳
	expiredAt := time.Now().Add(config.GetDuration("jwt.ttl") * time.Second).Unix()

	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    config.GetString("app.name"),
		},
	}

	// 生成令牌结构体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 将令牌结构提进行加密
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
