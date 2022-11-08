package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Params
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Validate Params with admin
		if username == "admin" && password == "admin" {
			// Create Token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set Claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "admin"
			claims["admin"] = true
			claims["iat"] = time.Now().Unix()
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}

			// Return token
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		m := "Hello " + name + "!"

		return c.JSON(http.StatusOK, map[string]string{
			"message": m,
		})
	}
}
