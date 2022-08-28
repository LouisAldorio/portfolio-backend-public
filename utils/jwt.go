package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateToken(username string) (string, error) {
	var jwtKey = []byte(os.Getenv("KEY"))
	var signingMethod = jwt.SigningMethodHS256
	var expiredTime = time.Now().AddDate(0, 1, 0).UnixNano() / int64(time.Millisecond)

	customClaim := UserClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(signingMethod, customClaim)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", gqlerror.Errorf(fmt.Sprintf("%s", err))
	}

	return signedToken, nil
}

func ValidateToken(t string) (*jwt.Token, error) {
	var jwtKey = []byte(os.Getenv("KEY"))
	token, _ := jwt.ParseWithClaims(t, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return jwtKey, nil
	})

	return token, nil
}
