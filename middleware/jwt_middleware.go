package middleware

import (
	"mini-project/helper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JwtMiddlewareBuyer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.UnauthorizedResponse(c)
		}

		tokenString := authHeader[len("Bearer "):]

		if tokenString == "" {
			return helper.UnauthorizedResponse(c)
		}

		jwtService := helper.NewJwtService()
		parsedToken, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedResponse(c)
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			role := claims["Role"].(string)

			if role != "buyer" {
				return helper.ForbiddenResponse(c)
			}
			if id, ok := claims["Id"].(float64); ok {
				c.Set("buyerId", int(id))
			} else {
				return helper.UnauthorizedResponse(c)
			}
		} else {
			return helper.UnauthorizedResponse(c)
		}

		return next(c)
	}
}

func JwtMiddlewareSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.UnauthorizedResponse(c)
		}

		tokenString := authHeader[len("Bearer "):]

		if tokenString == "" {
			return helper.UnauthorizedResponse(c)
		}

		jwtService := helper.NewJwtService()
		parsedToken, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedResponse(c)
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			role := claims["Role"].(string)

			if role != "seller" {
				return helper.ForbiddenResponse(c)
			}
			if id, ok := claims["Id"].(float64); ok {
				c.Set("sellerId", int(id))
			} else {
				return helper.UnauthorizedResponse(c)
			}
		} else {
			return helper.UnauthorizedResponse(c)
		}

		return next(c)
	}
}
