package utils

import (
	adminConf "cap-club/cmd/admin-service/config"
	userConf "cap-club/cmd/user-service/config"
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

type UserClaims struct {
	Username       string `json:"username"`
	StandardClaims jwt.StandardClaims
}

func (c *UserClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func GenerateJWTUser(username string, conf *userConf.ServiceConfig) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.JWTKey)) //
}

// --------

type AdminClaims struct {
	AdminName      string `json:"admin"`
	StandardClaims jwt.StandardClaims
}

func (c *AdminClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func GenerateJWTAdmin(adminName string, conf *adminConf.ServiceConfig) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &AdminClaims{
		AdminName: adminName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.JWTKey))
}
