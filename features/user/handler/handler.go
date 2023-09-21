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
	_, _, companyId := middlewares.ExtractTokenUserId(c)

	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	// userInput.CompanyID = uint(companyId)
	fmt.Println("USER INPUT", userInput)
	fmt.Println("USER INPUT", userInput.CompanyID, userInput.LevelID, userInput.ManagerID)
	userCore := RequestToCore(*userInput)

	// fmt.Println("company id", companyId)
	err := handler.userService.Create(userCore, companyId)
	if err != nil {

		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else if strings.Contains(err.Error(), "' for key 'users.email'") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "User with this email already exists", nil))
		} else if strings.Contains(err.Error(), "Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`))") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Role with this id is not found", err.Error()))
		} else if strings.Contains(err.Error(), "Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_level` FOREIGN KEY (`level_id`) REFERENCES `employee_levels` (`id`))") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusConflict, "Level with this id is not found", err.Error()))
		} else if strings.Contains(err.Error(), "Cannot add or update a child row: a foreign key constraint fails (`hris_kelompok1_2`.`users`, CONSTRAINT `fk_users_company` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`))") {
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
	fmt.Println("CMPNY", dataLogin.Company.ID)
	response := LoginResponse{
		ID:          dataLogin.ID,
		Email:       dataLogin.Email,
		RoleName:    dataLogin.Role.RoleName,
		CompanyName: dataLogin.Company.CompanyName,
		Level:       dataLogin.Level.Level,
		Fullname:    dataLogin.Fullame,
		CompanyID:   dataLogin.Company.ID,
		Token:       token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success login", response))
}

func (handler *UserHandler) GetProfileUser(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractTokenUserId(c)

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
		CreatedAt:       result.CreatedAt,
		UpdatedAt:       result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", usersResponse))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))
	_, _, companyId := middlewares.ExtractTokenUserId(c)

	managerId, _ := strconv.Atoi(c.QueryParam("manager_id"))

	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 1000
	}

	result, err := handler.userService.GetAll(int(pageNumber), int(pageSize), managerId, companyId)
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
	idToken, _, _ := middlewares.ExtractTokenUserId(c)
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

func (handler *UserHandler) UpdateOtherProfile(c echo.Context) error {
	_, roleName, _ := middlewares.ExtractTokenUserId(c)

	if (roleName != "Superadmin") && (roleName != "HR") {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Forbidden Access You are not Superadmin or HR", nil))
	} else {
		userInput := new(UpdateProfileRequest)
		errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
		}

		id := c.Param("user_id")
		idConv, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
		}

		userInput.UserId = uint(idConv)

		userProfileCore := RequestUpdateProfileToCore(*userInput)
		err := handler.userService.UpdateOtherProfile(idConv, userProfileCore)
		// fmt.Println("ERROR", err)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
		}
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Update data", nil))
	}

}

func (handler *UserHandler) GetOtherProfileUser(c echo.Context) error {
	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	_, _, companyId := middlewares.ExtractTokenUserId(c)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}

	result, err := handler.userService.SelectOtherProfile(idConv)
	fmt.Println("Result company ID ", result.Company.ID)
	fmt.Println("Result company ID ", companyId)
	fmt.Println(companyId != int(result.Company.ID))
	if companyId != int(result.Company.ID) {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusNotFound, "You cannot see profile from user different company with you", nil))
	} else {
		if err != nil {

			if strings.Contains(err.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
			} else if strings.Contains(err.Error(), "record not found") {
				return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusNotFound, "User with this id is not found", nil))
			} else {
				return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
			}

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
			CreatedAt:       result.CreatedAt,
			UpdatedAt:       result.UpdatedAt,
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read detail data", usersResponse))
	}
}

func (handler *UserHandler) DeleteOtherProfile(c echo.Context) error {
	id := c.Param("user_id")
	idUser, errConv := strconv.Atoi(id)

	_, roleName, _ := middlewares.ExtractTokenUserId(c)

	if (roleName != "Superadmin") && (roleName != "HR") {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "Forbidden Access You are not Superadmin or HR", nil))
	} else {
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error id not valid", nil))
		}

		err := handler.userService.DeleteOtherProfile(uint(idUser))
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
			} else {
				return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))

			}
		}
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete user", nil))
	}

}
