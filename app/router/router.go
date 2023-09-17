package router

import (
	"group-project-3/app/middlewares"
	_userData "group-project-3/features/user/data"
	_userHandler "group-project-3/features/user/handler"
	_userService "group-project-3/features/user/service"

	_roleData "group-project-3/features/role/data"
	_roleHandler "group-project-3/features/role/handler"
	_roleService "group-project-3/features/role/service"

	_companyData "group-project-3/features/company/data"
	_companyHandler "group-project-3/features/company/handler"
	_companyService "group-project-3/features/company/service"

	_userDetailData "group-project-3/features/userDetail/data"
	_userDetailHandler "group-project-3/features/userDetail/handler"
	_userDetailService "group-project-3/features/userDetail/service"

	_employeeLevelData "group-project-3/features/employeeLevel/data"
	_employeeLevelHandler "group-project-3/features/employeeLevel/handler"
	_employeeLevelService "group-project-3/features/employeeLevel/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	roleData := _roleData.New(db)
	roleService := _roleService.New(roleData)
	roleHandlerAPI := _roleHandler.New(roleService)

	companyData := _companyData.New(db)
	companyService := _companyService.New(companyData)
	companyHandlerAPI := _companyHandler.New(companyService)

	userDetailData := _userDetailData.New(db)
	userDetailService := _userDetailService.New(userDetailData)
	userDetailHandlerAPI := _userDetailHandler.New(userDetailService)

	employeeLevelData := _employeeLevelData.New(db)
	employeeLevelService := _employeeLevelService.New(employeeLevelData)
	employeeLevelHandlerAPI := _employeeLevelHandler.New(employeeLevelService)

	//Authentikasi
	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/my-profile", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.PUT("/my-profile", userHandlerAPI.UpdateMyProfile, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.GetAllUser)

	e.POST("/roles", roleHandlerAPI.CreateRole, middlewares.JWTMiddleware())
	e.GET("/roles", roleHandlerAPI.GetAllRole, middlewares.JWTMiddleware())
	e.PUT("/roles/:role_id", roleHandlerAPI.UpdateRole, middlewares.JWTMiddleware())
	e.DELETE("/roles/:role_id", roleHandlerAPI.DeleteRole, middlewares.JWTMiddleware())

	e.POST("/companies", companyHandlerAPI.CreateCompany, middlewares.JWTMiddleware())
	e.DELETE("/companies/:company_id", companyHandlerAPI.DeleteCompany, middlewares.JWTMiddleware())
	e.GET("/companies", companyHandlerAPI.GetAllCompany, middlewares.JWTMiddleware())
	e.GET("/companies/:company_id", companyHandlerAPI.GetCompanyId, middlewares.JWTMiddleware())
	e.PUT("/companies/:company_id", companyHandlerAPI.UpdateById, middlewares.JWTMiddleware())

	e.PUT("/user-details/:user_id", userDetailHandlerAPI.Update, middlewares.JWTMiddleware())
	e.GET("/user-details/:user_id", userDetailHandlerAPI.FindById, middlewares.JWTMiddleware())
	e.DELETE("/user-details/:user_id", userDetailHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.POST("/user-details", userDetailHandlerAPI.Create, middlewares.JWTMiddleware())

	e.PUT("/employee-levels/:id", employeeLevelHandlerAPI.Update, middlewares.JWTMiddleware())
	e.GET("/employee-levels/:id", employeeLevelHandlerAPI.FindById, middlewares.JWTMiddleware())
	e.DELETE("/employee-levels/:id", employeeLevelHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.POST("/employee-levels", employeeLevelHandlerAPI.Create, middlewares.JWTMiddleware())
	e.GET("/employee-levels", employeeLevelHandlerAPI.FindAll, middlewares.JWTMiddleware())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"===message": "Welcome to kelompok 1 Group Project3==",
		})
	})
}
