package handler

import (
	"fmt"
	"group-project-3/features/role"
	"group-project-3/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService role.RoleServiceInterface
}

func New(service role.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{
		roleService: service,
	}
}

func (handler *RoleHandler) CreateRole(c echo.Context) error {
	roleInput := new(RoleRequest)
	fmt.Println("ROLE INPUT", &roleInput)
	errBind := c.Bind(&roleInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	userCore := RequestToCore(*roleInput)
	err := handler.roleService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else if strings.Contains(err.Error(), "' for key 'users.email'") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "User with this email already exists", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *RoleHandler) GetAllRole(c echo.Context) error {
	result, err := handler.roleService.GetAll()

	fmt.Println("RESULT ROLE", result)

	var roleResponse []RoleResponse

	for _, value := range result {
		roleResponse = append(roleResponse, RoleResponse{
			ID:       value.ID,
			RoleName: value.RoleName,
		})
	}
	if err != nil {

		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read all roles", roleResponse))
}

func (handler *RoleHandler) UpdateRole(c echo.Context) error {
	id := c.Param("role_id")
	idRole, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	roleInput := new(RoleRequest)

	errBind := c.Bind(&roleInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	roleCore := RequestToCore(*roleInput)
	err := handler.roleService.UpdateRole(uint(idRole), roleCore)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusNotFound, err.Error(), nil))

		}
	} else {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Update Role", nil))
	}

}

func (handler *RoleHandler) DeleteRole(c echo.Context) error {
	id := c.Param("role_id")

	idRole, errConv := strconv.Atoi(id)
	fmt.Println("ID Role", idRole)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	err := handler.roleService.DeleteRole(idRole)
	fmt.Println("DEBUG", err)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Delete Role", nil))
}
