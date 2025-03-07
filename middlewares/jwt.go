package middlewares

import (
	"errors"
	"log"
	"my-project-be/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET)})
}

func GenerateJWT(id uint, nama string) (string, error) {
	var data = jwt.MapClaims{}
	data["id"] = id
	data["nama"] = nama
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()
	processToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := processToken.SignedString([]byte(config.JWTSECRET))
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)
			}
		}()
		return "", errors.New("error jwt creation")
	}

	return result, nil
}
	
