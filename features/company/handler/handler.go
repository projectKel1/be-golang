package handler

import (
	"fmt"
	"group-project-3/app/middlewares"
	"group-project-3/features/company"
	"group-project-3/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type companyHandler struct {
	companyService company.CompanyServiceInterface
}

func New(service company.CompanyServiceInterface) *companyHandler {
	return &companyHandler{
		companyService: service,
	}
}

func (handler *companyHandler) CreateCompany(c echo.Context) error {
	userInput := new(CompanyRequest)
	errBind := c.Bind(&userInput)
	_, roleName, _ := middlewares.ExtractTokenUserId(c)
	if roleName != "Superadmin" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access denied", nil))
	}
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	companyCore := RequestToCore(*userInput)
	err := handler.companyService.Create(companyCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else if strings.Contains(err.Error(), "' for key 'company.name'") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Company with this name already exists", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *companyHandler) DeleteCompany(c echo.Context) error {
	id := c.Param("company_id")
	idCompany, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}

	err := handler.companyService.DeleteById(uint(idCompany))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete company", nil))
}

func (handler *companyHandler) GetAllCompany(c echo.Context) error {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))

	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	result, err := handler.companyService.GetAll(int(pageNumber), int(pageSize))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var companyResponse []CompanyResponse
	for _, value := range result {
		companyResponse = append(companyResponse, CompanyResponse{
			ID:          value.ID,
			Name:        value.Name,
			Address:     value.Address,
			Description: value.Description,
			Email:       value.Email,
			Type:        value.Type,
			Image:       value.Image,
			Visi:        value.Visi,
			Misi:        value.Misi,
			StartedHour: value.StartedHour,
			EndedHour:   value.EndedHour,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", companyResponse))
}

func (handler *companyHandler) GetCompanyId(c echo.Context) error {
	id := c.Param("company_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}

	result, err := handler.companyService.GetById(uint(idConv))
	if err != nil {

		return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "data not found", nil))
	}
	// mapping dari struct core to struct response
	resultResponse := CompanyResponse{
		ID:          result.ID,
		Name:        result.Name,
		Address:     result.Address,
		Description: result.Description,
		Email:       result.Email,
		Type:        result.Type,
		Image:       result.Image,
		Visi:        result.Visi,
		Misi:        result.Misi,
		StartedHour: result.StartedHour,
		EndedHour:   result.EndedHour,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}

func (handler *companyHandler) UpdateById(c echo.Context) error {
	userInput := new(CompanyRequest)

	id := c.Param("company_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error data id. data not valid", nil))
	}
	fmt.Println("COMPANY INPUT", userInput)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	companyCore := RequestToCore(*userInput)
	err := handler.companyService.EditById(uint(idParam), companyCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update data", nil))
}
