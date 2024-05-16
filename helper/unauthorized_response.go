package helper

import (
	dto_base "mini-project/dto/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UnauthorizedResponse(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, dto_base.BaseResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Data:    nil,
	})
}

func ForbiddenResponse(c echo.Context) error {
	return c.JSON(http.StatusForbidden, dto_base.BaseResponse{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Data:    nil,
	})
}
