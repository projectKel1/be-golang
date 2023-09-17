package handler

import (
	"errors"
	"group-project-3/exception"
	"group-project-3/features/employeeLevel"
	"group-project-3/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeLevelHandlerImpl struct {
	Service employeeLevel.EmployeeLevelServiceInterface
}

func New(service employeeLevel.EmployeeLevelServiceInterface) employeeLevel.EmployeeLevelHandlerInterface {
	return &EmployeeLevelHandlerImpl{Service: service}
}

// FindAll implements employeeLevel.EmployeeLevelHandlerInterface
func (handler *EmployeeLevelHandlerImpl) FindAll(c echo.Context) error {
	results, err := handler.Service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	responses := SliceEmployeeLevelToSliceResponse(results)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", responses))
}

// Create implements employeeLevel.EmployeeLevelHandlerInterface
func (handler *EmployeeLevelHandlerImpl) Create(c echo.Context) error {
	var dataRequest EmployeeLevelRequest
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

// Delete implements employeeLevel.EmployeeLevelHandlerInterface
func (handler *EmployeeLevelHandlerImpl) Delete(c echo.Context) error {
	userId := c.Param("id")
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

// FindById implements employeeLevel.EmployeeLevelHandlerInterface
func (handler *EmployeeLevelHandlerImpl) FindById(c echo.Context) error {
	userId := c.Param("id")
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
	response := EmployeeLevelToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success", response))
}

// Update implements employeeLevel.EmployeeLevelHandlerInterface
func (handler *EmployeeLevelHandlerImpl) Update(c echo.Context) error {
	var dataRequest EmployeeLevelRequest
	err := c.Bind(&dataRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, exception.ErrBadRequest.Error(), nil))
	}

	dataRequest.ID = uint(intId)

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
