package router

import (
	"group-project-3/app/middlewares"
	_userData "group-project-3/features/user/data"
	_userHandler "group-project-3/features/user/handler"
	_userService "group-project-3/features/user/service"

	_roleData "group-project-3/features/role/data"
	_roleHandler "group-project-3/features/role/handler"
	_roleService "group-project-3/features/role/service"
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

	//Authentikasi
	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/my-profile", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())

	e.POST("/roles", roleHandlerAPI.CreateRole, middlewares.JWTMiddleware())
	e.GET("/roles", roleHandlerAPI.GetAllRole, middlewares.JWTMiddleware())
	e.PUT("/roles/:role_id", roleHandlerAPI.UpdateRole, middlewares.JWTMiddleware())
	e.DELETE("/roles/:role_id", roleHandlerAPI.DeleteRole, middlewares.JWTMiddleware())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"===message": "Welcome to kelompok 1 Group Project3==",
		})
	})
}
