package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// UserLoader represents middleware that require user to be loggedin.
type UserLoader struct{}

func (ul *UserLoader) Do(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			return c.String(http.StatusUnauthorized, "Malformed token")
		}

		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Bad signing method: %v", token.Header["alg"])
			}
			return []byte("some-secret"), nil
		})
		_ = err

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			// ctx := context.WithValue(c.Request().Context(), "props", claims)
			//Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.Mapclaim)
			c.Set("props", claims)
			next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
}
