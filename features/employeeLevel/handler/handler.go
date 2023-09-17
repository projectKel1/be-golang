package handler

import (
	"fmt"
	"group-project-3/features/employeeLevel"
	"group-project-3/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type EmployeeLevelHandler struct {
	employeeLevelService employeeLevel.EmployeeLevelServiceInterface
}

func New(service employeeLevel.EmployeeLevelServiceInterface) *EmployeeLevelHandler {
	return &EmployeeLevelHandler{
		employeeLevelService: service,
	}
}

func (handler *EmployeeLevelHandler) CreateEmployeeLevel(c echo.Context) error {
	employeeLevelInput := new(EmployeeLevelRequest)
	fmt.Println("Employee Level INPUT", &employeeLevelInput)
	errBind := c.Bind(&employeeLevelInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	employeeLevelCore := RequestToCore(*employeeLevelInput)
	err := handler.employeeLevelService.Create(employeeLevelCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *EmployeeLevelHandler) GetAllEmployeeLevel(c echo.Context) error {
	result, err := handler.employeeLevelService.GetAll()

	fmt.Println("RESULT ROLE", result)

	var employeeLevelResponse []EmployeLevelResponse

	for _, value := range result {
		employeeLevelResponse = append(employeeLevelResponse, EmployeLevelResponse{
			ID:    value.ID,
			Level: value.Level,
		})
	}
	if err != nil {

		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read all employee level", employeeLevelResponse))
}

func (handler *EmployeeLevelHandler) UpdateEmployeeLevel(c echo.Context) error {
	id := c.Param("level_id")
	idLevel, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	employeeLevelInput := new(EmployeeLevelRequest)

	errBind := c.Bind(&employeeLevelInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	employeeLevelCore := RequestToCore(*employeeLevelInput)
	err := handler.employeeLevelService.UpdateEmployeeLevel(uint(idLevel), employeeLevelCore)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusNotFound, err.Error(), nil))

		}
	} else {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Update Employee level", nil))
	}

}

func (handler *EmployeeLevelHandler) DeleteEmployeeLevel(c echo.Context) error {
	id := c.Param("level_id")

	idEmployeeLevel, errConv := strconv.Atoi(id)
	fmt.Println("ID Employee Level", idEmployeeLevel)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	err := handler.employeeLevelService.DeleteEmployeeLevel(idEmployeeLevel)
	fmt.Println("DEBUG", err)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Delete Employee Level", nil))
}
