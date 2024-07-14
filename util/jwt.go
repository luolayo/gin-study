package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"time"
)

type JwtCustomClaims struct {
	ID       int
	Username string
	jwt.RegisteredClaims
}

func CreateToken(user model.User) (string, error) {
	var jwtSecret = []byte(global.SysConfig.JwtSecret)
	iJwtCustomClaims := JwtCustomClaims{
		ID:       int(user.ID),
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (JwtCustomClaims, error) {
	var jwtSecret = []byte(global.SysConfig.JwtSecret)
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
			return *claims, nil
		}
	}
	return iJwtCustomClaims, err
}
