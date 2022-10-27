package jwt

import (
	"be_project3team3/config"
	"errors"
	"log"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func JWTMiddleware() echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(config.JWT_SECRET),
	})

}

func GenerateJWTToken(id uint, shopname string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	claims["shop_name"] = shopname
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("Error generate JWT Token. error ", err)
		return "", errors.New("Error generate JWT Token.")
	}

	return str, nil
}

func ExtractIdToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id_user := claims["id_user"].(float64)
		return uint(id_user)
	}

	return 0
}

func ExtractToken(c echo.Context) (uint, int64) {
	token := c.Get("user").(*jwt.Token)
	// log.Println("\n\n\nisi token\n", token, "\n\n")
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return uint(claims["id"].(float64)), int64(claims["exp"].(float64))
	}

	return 0, 0
}

func ExtracShopNameToken(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return claims["shop_name"].(string)
	}

	return ""
}

func IsAuthorized(c echo.Context) error {
	id, exp := ExtractToken(c)

	if id == 0 {
		return errors.New("Request not authorized. Please check token. User not found.")
	} else if time.Now().Unix() > exp {
		return errors.New("Request not authorized. Please check token. Expired token.")
	} else {
		return nil
	}
}
