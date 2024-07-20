package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"time"
)

type JwtCustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func CreateToken(user model.User) (string, error) {
	var jwtSecret = []byte(global.SysConfig.JwtSecret)
	iJwtCustomClaims := JwtCustomClaims{
		ID:   int(user.Uid),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		global.LOG.Error("Failed to create token %v", err)
		return "", err
	}
	err = global.Redis.Set(user.Name, tokenString, 24*time.Hour)
	if err != nil {
		global.LOG.Error("Failed to set redis %v", err)
		return "", err
	}
	return tokenString, nil
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

func UpdateToken(claims JwtCustomClaims) (string, error) {
	err := global.Redis.Del(claims.Name)
	if err != nil {
		return "", err
	}
	token, err := CreateToken(model.User{Uid: uint(claims.ID), Name: claims.Name})
	if err != nil {
		return "", err
	}
	err = global.Redis.Set(claims.Name, token, 24*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}
