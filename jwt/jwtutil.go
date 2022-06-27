package jwt

import (
	jwt "github.com/golang-jwt/jwt"
)

type Jwt struct {
	Name string
}

func NewJwt(name string) *Jwt {
	return &Jwt{name}
}

func (j *Jwt) GetToken(name string) (string, error) {
	signingKey := j.GenKey(name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func (j *Jwt) VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := j.GetKey(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
