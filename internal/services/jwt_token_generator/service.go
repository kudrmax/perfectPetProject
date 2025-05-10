package jwt_token_generator

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	UnexpectedSigningMethodErr = errors.New("unexpected signing method")
	InvalidTokenErr            = errors.New("invalid token")
)

type Service struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewService(secret string, duration time.Duration) *Service {
	return &Service{
		secretKey:     secret,
		tokenDuration: duration,
	}
}

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *Service) GenerateToken(userId int) (string, error) {
	expirationTime := time.Now().Add(s.tokenDuration)

	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *Service) ParseToken(tokenStr string) (userId int, err error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, UnexpectedSigningMethodErr
		}
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, InvalidTokenErr
	}

	return claims.UserId, nil
}
