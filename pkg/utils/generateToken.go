package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret []byte
var RefreshTokenSecret []byte

func InitializeJWTSecret(secret string, refreshSecret string) {
	jwtSecret = []byte(secret)
	RefreshTokenSecret = []byte(refreshSecret)
}

func GenerateAccessToken(userID uint, email string) (string, error) {
	location, _ := time.LoadLocation("Africa/Ouagadougou")
	expireAt := time.Now().Add(15 * time.Minute).In(location)
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now().In(location)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claim, nil
	}

	return nil, errors.New("token invalide")
}

func GenerateRefreshToken(userID uint, email string) (string, *time.Time, error) {
	location, _ := time.LoadLocation("Africa/Ouagadougou")

	expireAt := time.Now().Add(30 * 24 * time.Hour).In(location)

	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now().In(location)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(RefreshTokenSecret)

	return tokenStr, &expireAt, err
}
