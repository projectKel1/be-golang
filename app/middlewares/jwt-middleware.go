package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte("sup3rs3cr3t"),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId uint, roleName string, level string, companyName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role_name"] = roleName
	claims["level"] = level
	claims["company_name"] = companyName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("sup3rs3cr3t"))

}

func ExtractTokenUserId(e echo.Context) (int, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		roleName := claims["role_name"].(string)
		return int(userId), roleName
	}
	return 0, ""
}
