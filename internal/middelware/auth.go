package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/go-api/internal/config"
)

type contextKey string

const UserKey contextKey = "userClaims"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Envs.AuthSecret), nil
		})
		// attempt to return better errors to the client
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed):
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
			case errors.Is(err, jwt.ErrTokenSignatureInvalid):
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
			case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
				return echo.NewHTTPError(http.StatusUnauthorized, "token expired or not active yet")
			default:
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
		}

		// flat decline if the token is invalid
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}

		// convert roles array to []interface
		roles := make([]string, len(claims["roles"].([]interface{})))
		for i, role := range claims["roles"].([]interface{}) {
			roles[i] = role.(string)
		}

		// ctx := context.WithValue(c.Request().Context(), UserKey, userClaims)
		// c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
