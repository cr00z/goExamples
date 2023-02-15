package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey}
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (m *JWTManager) Generate(user *User) (string, error) {
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.tokenDuration).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.secretKey)
}

func (m *JWTManager) Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("token claims are not in type")
	}

	return claims, nil
}