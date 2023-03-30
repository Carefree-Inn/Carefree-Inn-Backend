package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

type Claims struct {
	Account string
	jwt.StandardClaims
}

func GenerateToken(account string) (string, error) {
	claims := Claims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    "CCNU-Inn",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("CCNU-Inn"))
	
	return token, errors.WithStack(err)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("CCNU-Inn"), nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			return claims, nil
		}
	}
	return nil, errors.WithStack(err)
}
