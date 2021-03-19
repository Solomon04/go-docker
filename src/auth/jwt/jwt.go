package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte(os.Getenv("APP_JWT_TOKEN"))
)

// GenerateToken generates a jwt token and assign a email to it's claims and return it
func GenerateToken(email string, uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["email"] = email
	claims["uuid"] = uuid
	claims["createdAt"] = time.Now().Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Println("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the email in it's claims
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, err
	}
}
