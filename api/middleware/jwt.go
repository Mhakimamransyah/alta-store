package middleware

import (
	"os"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
	})
}

func ExtractTokenKey(c echo.Context, key string) interface{} {
	admin := c.Get("user").(*jwt.Token)
	if admin.Valid {
		claims := admin.Claims.(jwt.MapClaims)
		return claims[key]
	}
	return ""
}
