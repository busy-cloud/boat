package web

import (
	"github.com/busy-cloud/boat/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Admin bool `json:"admin,omitempty"`
}

var JwtKey = "boat"
var JwtExpire = time.Hour * 24 * 30

func JwtGenerate(id string, admin bool) (string, error) {
	var claims Claims
	claims.ID = id
	claims.Admin = admin
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(JwtExpire))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func JwtVerify(str string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(str, &claims, func(token *jwt.Token) (any, error) {
		return config.GetString(MODULE, "jwt_key"), nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
