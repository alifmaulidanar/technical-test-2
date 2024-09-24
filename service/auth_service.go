package service

import (
	"os"
	"technical-test-II/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// membuat token JWT untuk user login
func GenerateJWT(user domain.User) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id":      user.UserID,
		"phone_number": user.PhoneNumber,
		"exp":          time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

// pengeeckan pin yang di-hash untuk keamanan
func CheckPinHash(pin, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin))
	return err == nil
}
