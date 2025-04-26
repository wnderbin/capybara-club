package utils

import (
	"cap-club/admin-service/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func CheckHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Claims struct {
	AdminName      string `json:"admin"`
	StandardClaims jwt.StandardClaims
}

func (c *Claims) Valid() error {
	return c.StandardClaims.Valid()
}

func GenerateJWT(adminName string, conf *config.ServiceConfig) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		AdminName: adminName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	return token.SignedString([]byte(conf.JWTKey))
}
