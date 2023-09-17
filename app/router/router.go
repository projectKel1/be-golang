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

	employeeLevelData := _employeeLevelData.New(db)
	employeeLevelService := _employeeLevelService.New(employeeLevelData)
	employeeLevelHandlerAPI := _employeeLevelHandler.New(employeeLevelService)

	//Authentikasi
	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/my-profile", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.GetAllUser)

	e.POST("/roles", roleHandlerAPI.CreateRole)
	e.GET("/roles", roleHandlerAPI.GetAllRole)
	e.PUT("/roles/:role_id", roleHandlerAPI.UpdateRole, middlewares.JWTMiddleware())
	e.DELETE("/roles/:role_id", roleHandlerAPI.DeleteRole, middlewares.JWTMiddleware())

	e.POST("/employee-level", employeeLevelHandlerAPI.CreateEmployeeLevel)
	e.GET("/employee-level", employeeLevelHandlerAPI.GetAllEmployeeLevel)
	e.PUT("/employee-level/:level_id", employeeLevelHandlerAPI.UpdateEmployeeLevel)
	e.DELETE("/employee-level/:level_id", employeeLevelHandlerAPI.DeleteEmployeeLevel)

	e.POST("/companies", companyHandlerAPI.CreateCompany)
	e.DELETE("/companies/:company_id", companyHandlerAPI.DeleteCompany, middlewares.JWTMiddleware())
	e.GET("/companies", companyHandlerAPI.GetAllCompany)
	e.GET("/companies/:company_id", companyHandlerAPI.GetCompanyId, middlewares.JWTMiddleware())
	e.PUT("/companies/:company_id", companyHandlerAPI.UpdateById, middlewares.JWTMiddleware())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"===message": "Welcome to kelompok 1 Group Project3==",
		})
	})
}
