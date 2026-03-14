package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64 
	Email string 
	jwt.RegisteredClaims
}

var jwtSecret = []byte("my-secret")

func GenerateJWT(c Claims) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*Claims,error){
	token,err := jwt.ParseWithClaims(tokenString,&Claims{},func(t *jwt.Token) (interface{},error){
		if _ , ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret,nil 
	})
	if err != nil {
		return nil, err 
	}

	claims,ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims,nil 
}