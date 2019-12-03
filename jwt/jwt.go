package jwt

import (
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
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("12134"))

	if err != nil {
		log.Fatal(nil)
	}

	return tokenString
}

func validating(jwtStr string) (jwt.MapClaims, Errs) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("12134"), nil
	})

	if err != nil {
		if token == nil {
			return nil, Errs{
				Message: "Token Format error",
				Code:    103,
			}
		} else {
			if !token.Valid {
				ve, ok := err.(*jwt.ValidationError)
				if ok && ve.Errors&jwt.ValidationErrorExpired != 0 {

					return nil, Errs{
						Message: "Token Expired",
						Code:    101,
					}
				}
				return nil, Errs{
					Message: "Token Invalid",
					Code:    102,
				}
			}
		}
	}

	return token.Claims.(jwt.MapClaims), Errs{}
}

type Errs struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
