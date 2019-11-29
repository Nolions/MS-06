package jwt

import (
	"fmt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"log"
	"time"
)

type user struct {
	id      int
	account string
}

type tokenSecret []byte

func signing(u *user) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      u.id,
		"account": u.account,
		"exp":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("123434"))

	if err != nil {
		log.Fatal(nil)
	}

	return tokenString
}

func validating(jwtStr string) (bool, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("123434"), nil
	})

	if err != nil {
		return false, err
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	return true, nil
}
