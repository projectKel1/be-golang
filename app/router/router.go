package router

import (
	"group-project-3/app/middlewares"
	_userData "group-project-3/features/user/data"
	_userHandler "group-project-3/features/user/handler"
	_userService "group-project-3/features/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	//Authentikasi
	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/my-profile", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Welcome To Kelompok 1",
		})
	})
}
