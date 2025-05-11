package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func NewGenerateJwt(username string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		NewLoggerConfigure("env.log", logrus.InfoLevel, err.Error(), logrus.InfoLevel)
		return err.Error(), err
	}

	secretKey := []byte(os.Getenv("ACCESS_SECRET"))

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		NewLoggerConfigure("env.log", logrus.InfoLevel, err.Error(), logrus.InfoLevel)
		return err.Error(), err
	}

	return signedToken, nil

}
