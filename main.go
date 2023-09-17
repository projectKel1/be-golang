package main

import (
	"group-project-3/app/config"
	"group-project-3/app/database"
	"group-project-3/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbSql := database.InitDBMysql(cfg)

	database.InitialMigration(dbSql)

	// create a new echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":80"))
}
