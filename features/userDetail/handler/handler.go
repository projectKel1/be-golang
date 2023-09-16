package handler

import (
	"errors"
	"group-project-3/exception"
	"group-project-3/features/userDetail"
	"group-project-3/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDetailHandlerImpl struct {
	Service userDetail.UserDetailServiceInterface
}

func New(service userDetail.UserDetailServiceInterface) userDetail.UserDetailHandlerInterface {
	return &UserDetailHandlerImpl{Service: service}
}

// Create implements userDetail.UserDetailHandlerInterface
func (handler *UserDetailHandlerImpl) Create(c echo.Context) error {
	var dataRequest UserDetailRequest
	err := c.Bind(&dataRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}
	dataEntity := RequestToEntity(dataRequest)
	err = handler.Service.Create(dataEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, exception.ErrInternalServer.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "created", nil))
}

// Delete implements userDetail.UserDetailHandlerInterface
func (handler *UserDetailHandlerImpl) Delete(c echo.Context) error {
	userId := c.Param("user_id")
	intId, err := strconv.Atoi(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}
	err = handler.Service.Delete(uint(intId))
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, exception.ErrIdIsNotFound.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, exception.ErrInternalServer.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", nil))
}

// FindById implements userDetail.UserDetailHandlerInterface
func (handler *UserDetailHandlerImpl) FindById(c echo.Context) error {
	userId := c.Param("user_id")
	intId, err := strconv.Atoi(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}
	result, err := handler.Service.FindById(uint(intId))
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, exception.ErrIdIsNotFound.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, exception.ErrInternalServer.Error(), nil))
		}
	}
	response := UserDetailToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", response))
}

// Update implements userDetail.UserDetailHandlerInterface
func (handler *UserDetailHandlerImpl) Update(c echo.Context) error {
	var dataRequest UserDetailRequest
	err := c.Bind(&dataRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}

	userId := c.Param("user_id")
	intId, err := strconv.Atoi(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}

	dataRequest.UserID = uint(intId)

	dataEntity := RequestToEntity(dataRequest)
	err = handler.Service.Update(dataEntity)

	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, exception.ErrIdIsNotFound.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, exception.ErrInternalServer.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", nil))

}
