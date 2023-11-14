package utils

import (
	"discordrm/api/config"
	"discordrm/api/interval/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var cfg *config.Config

func NewJWT(_cfg *config.Config) {
	cfg = _cfg
}

func CreateToken(user models.User, expiredIn int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UID,
	})

	// set ttl
	token.Claims.(jwt.MapClaims)["exp"] = expiredIn

	strToken, err := token.SignedString([]byte(cfg.JWT.SecretKey))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file": "utils/jwt.go",
			"func": "CreateToken",
		}).Error(err)
		return "", err
	}

	return strToken, nil
}

func ValidateToken(strToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.SecretKey), nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, err
	}

	return claims, nil
}
