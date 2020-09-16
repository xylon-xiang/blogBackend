package util

import (
	"blogBackend/config"
	"blogBackend/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"time"
)

type JwtClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

var JwtConfig middleware.JWTConfig

func GenerateJwtToken(id string) (*model.UserLogReturnModule, error) {

	var logReturn model.UserLogReturnModule

	claims := JwtClaim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().
				Add(time.Duration(config.Config.Jwt.TokenLife) * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(config.Config.Jwt.SigningSecret))
	if err != nil {
		return nil, err
	}

	logReturn.JwtToken = token

	return &logReturn, nil

}

func init() {

	JwtConfig.SigningKey = []byte(config.Config.Jwt.SigningSecret)
	JwtConfig.Claims = &JwtClaim{}
	JwtConfig.ContextKey = config.Config.Jwt.ContentKey

}
