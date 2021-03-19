package hasher

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
)

// Make makes a bcrypt hash from a given value
func Make(password string) string {
	bcryptRounds, _ := strconv.Atoi(os.Getenv("BCRYPT_ROUNDS"))
	pass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcryptRounds)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// Check checks whether a hash and its password are equivalent under bcrypt
func Check(hashedPassword string, password string) bool {
	hash := []byte(hashedPassword)
	pass := []byte(password)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false
	} else {
		return true
	}
}
