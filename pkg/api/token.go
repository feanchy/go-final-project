package api

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "secret"

func createHash(password string) string {
	hash := sha256.Sum256([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func createToken(password string) (string, error) {

	claims := jwt.MapClaims{
		"hash": createHash(password),
		"exp":  time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(jwtSecret))
}
