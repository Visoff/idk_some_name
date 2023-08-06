package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Controller struct {
	SecretKey []byte
}

func New(SecretKey []byte) *Controller {
	return &Controller{SecretKey}
}

func Use() *Controller {
	return New([]byte("Secret key"))
}

func (c *Controller) Encode(data interface{}) (string, error) {
	claims := jwt.MapClaims{
		"data": data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(c.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c *Controller) Verify(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return c.SecretKey, nil
	})

	if err != nil {
		return nil, errors.New("token parse error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("JWT token is invalid")
	}

	return claims["data"], nil
}
