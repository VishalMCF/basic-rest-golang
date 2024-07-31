package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// GetValue get value returns configuration based on a given key from the .env file
func GetValue(key string) string {
	// load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// return the value based on the given key
	return os.Getenv(key)
}

func GenerateNewAccessToken() (string, error) {
	secret := GetValue("JWT_SECRET_KEY")
	minutesCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
