package handler

import (
	"fmt"
	"group-project-3/app/middlewares"
	"group-project-3/features/user"
	"group-project-3/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := new(UserRequest)
	fmt.Println("USER INPUT", &userInput)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	userCore := RequestToCore(*userInput)
	err := handler.userService.Create(userCore)
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

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	fmt.Println("ROLE", dataLogin)
	response := LoginResponse{
		ID:     dataLogin.ID,
		Email:  dataLogin.Email,
		RoleID: int(dataLogin.RoleID),
		Status: dataLogin.Status,
		Token:  token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success login", response))
}

func (handler *UserHandler) GetProfileUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	fmt.Println("id:", idToken)
	result, err := handler.userService.GetProfile(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	usersResponse := ProfileResponse{
		ID:              result.ID,
		Fullname:        result.Fullame,
		Email:           result.Email,
		Password:        result.Password,
		PhoneNumber:     result.PhoneNumber,
		Address:         result.Address,
		UrlPhoto:        result.UrlPhoto,
		Gender:          result.Gender,
		Status:          result.Status,
		RoleID:          int(result.RoleID),
		CompanyID:       int(result.CompanyId),
		NoNik:           result.NoNik,
		NoKK:            result.NoKK,
		NoBpjs:          result.NoBpjs,
		EmergencyName:   result.EmergencyName,
		EmergencyStatus: result.EmergencyStatus,
		EmergencyPhone:  result.EmergencyPhone,
		CreatedAt:       result.CreatedAt,
		UpdatedAt:       result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
}
