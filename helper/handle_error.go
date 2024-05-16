package helper

import (
	dto_base "mini-project/dto/base"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Data:    validationErr.Error(),
		})
	}

	if err.Error() == EmailAlreadyExist.Error() {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: EmailAlreadyExist.Error(),
			Data:    nil,
		})
	}

	if err.Error() == InvalidLogin.Error() {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: InvalidLogin.Error(),
			Data:    nil,
		})
	}

	if err.Error() == TypeTrashAlreadyExist.Error() {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: TypeTrashAlreadyExist.Error(),
			Data:    nil,
		})
	}

	if err.Error() == TrashAlreadyExist.Error() {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: TrashAlreadyExist.Error(),
			Data:    nil,
		})
	}

	if err.Error() == TypeTrashNotFound.Error() {
		return c.JSON(http.StatusNotFound, dto_base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: TypeTrashNotFound.Error(),
			Data:    nil,
		})
	}

	if err.Error() == DataNotFound.Error() {
		return c.JSON(http.StatusNotFound, dto_base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: DataNotFound.Error(),
			Data:    nil,
		})
	}

	if err.Error() == TrashNotFound.Error() {
		return c.JSON(http.StatusNotFound, dto_base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: TrashNotFound.Error(),
			Data:    nil,
		})
	}

	if err.Error() == ErrorIdParam.Error() {
		return c.JSON(http.StatusBadRequest, dto_base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: ErrorIdParam.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusInternalServerError, dto_base.BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err,
	})

}
