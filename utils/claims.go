package utils

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
}

var (
	ErrTokenExpired     error = errors.New("token is expired")
	ErrTokenNotValidYet error = errors.New("token not active yet")
	ErrTokenMalformed   error = errors.New("that's not even a token")
	ErrTokenInvalid     error = errors.New("couldn't handle this token")
)

func EncryptClaims(claims *Claims, key *rsa.PrivateKey) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return t.SignedString(key)
}

func DecryptToken(token string, key *rsa.PublicKey) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
