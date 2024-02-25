package helper

import (
	"fmt"
	"toko1/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID uint   `json:"user_id"`
}

var (
	APPLICATION_NAME   string        = "Toko Online"
	LOGIN_EXP_DURATION time.Duration = time.Duration(24) * time.Hour
)

func GenerateToken(user entity.User) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXP_DURATION)),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("KEY")

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, ReturnIfError(err)
}

func ValidateToken(tokenString string) (uint, error) {
	var (
		userID uint
	)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		secretKey := os.Getenv("KEY")

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = uint(claims["user_id"].(float64))
		return userID, err
	}
	return userID, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}