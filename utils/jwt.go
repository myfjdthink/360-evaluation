package utils

import (
	"360-evaluation/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtConfig = config.Config.JWT

type Claims struct {
	userId uint
	jwt.StandardClaims
}

func GenerateToken(userId uint) (string, error) {
	now := time.Now()
	expire := now.Add(jwtConfig.Expire)
	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "360",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(jwtConfig.Secret))
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
