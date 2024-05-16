package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type DataClaim struct {
	Role string
	Id   int
	jwt.StandardClaims
}

type JwtService interface {
	GenerateToken(id int, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() JwtService {
	return &jwtService{}
}

func (service *jwtService) GenerateToken(id int, role string) (string, error) {
	claims := DataClaim{
		role,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			Issuer:    "go-jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))
	return t, err
}

func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(secret), nil
	})
}
