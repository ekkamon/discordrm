package utils

import (
	"discordrm/api/config"
	"discordrm/api/interval/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var cfg *config.Config

func NewJWT(_cfg *config.Config) {
	cfg = _cfg
}

func CreateToken(user models.User) (string, int64, error) {
	expired := time.Now().Add(30 * (24 * time.Hour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":        user.UID,
		"updated_at": user.UpdatedAt,
		"expired":    expired,
	})

	strToken, err := token.SignedString([]byte(cfg.JWT.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return strToken, expired, nil
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
