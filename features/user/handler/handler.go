package handler

import (
	"fmt"
	"group-project-3/app/middlewares"
	"group-project-3/features/user"
	"group-project-3/helpers"
	"net/http"
	"strconv"
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
	fmt.Println("USER INPUT", userInput)
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
		} else if strings.Contains(err.Error(), "Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`))") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Role with this id is not found", err.Error()))
		} else if strings.Contains(err.Error(), "Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_level` FOREIGN KEY (`level_id`) REFERENCES `employee_levels` (`id`))") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Level with this id is not found", err.Error()))
		} else if strings.Contains(err.Error(), "Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_company` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`))") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Company with this id is not found", err.Error()))
		} else if userInput.LevelID == 0 {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Level Id is required", nil))
		} else if userInput.CompanyID == 0 {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Company Id is required", nil))
		} else if userInput.RoleID == 0 {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Role ID is required", nil))
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
		ID:          dataLogin.ID,
		Email:       dataLogin.Email,
		RoleName:    dataLogin.Role.RoleName,
		CompanyName: dataLogin.Company.CompanyName,
		Level:       dataLogin.Level.Level,
		Token:       token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success login", response))
}

func (handler *UserHandler) GetProfileUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

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
		UrlPhoto:        result.UrlPhoto,
		RoleName:        result.Role.RoleName,
		Level:           result.Level.Level,
		Status:          result.Status,
		CompanyName:     result.Company.CompanyName,
		NoNik:           result.UserDetail.Nik,
		PhoneNumber:     result.UserDetail.PhoneNumber,
		Gender:          result.UserDetail.Gender,
		Address:         result.UserDetail.Address,
		NoBpjs:          result.UserDetail.Bpjs,
		NoKK:            result.UserDetail.NoKK,
		Npwp:            result.UserDetail.Npwp,
		EmergencyName:   result.UserDetail.EmergencyName,
		EmergencyStatus: result.UserDetail.EmergencyStatus,
		EmergencyPhone:  result.UserDetail.EmergencyPhone,
		// RoleID: int(result.RoleID),
		// CompanyID:       int(result.CompanyId),

		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))

	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	result, err := handler.userService.GetAll(int(pageNumber), int(pageSize))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var userResponse []UserResponse
	for _, value := range result {

		userResponse = append(userResponse, UserResponse{
			ID:          value.ID,
			Fullname:    value.Fullame,
			Email:       value.Email,
			RoleName:    value.RoleName,
			UrlPhoto:    value.UrlPhoto,
			Status:      value.Status,
			LevelName:   value.LevelName,
			CompanyName: value.CompanyName,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read all users", userResponse))
}

func (handler *UserHandler) UpdateMyProfile(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	userInput := new(UpdateProfileRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	userInput.UserId = uint(idToken)

	userProfileCore := RequestUpdateProfileToCore(*userInput)
	err := handler.userService.UpdateProfile(idToken, userProfileCore)
	// fmt.Println("ERROR", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Update data", nil))
}
