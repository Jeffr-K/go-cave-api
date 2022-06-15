package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"microuser/routes"
)

//
func main() {

	e := echo.New()

	e.Debug = true

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// router
	userGroup := e.Group("/user")
	{
		userGroup.POST("", routes.CreateUserController)
		userGroup.DELETE("/:id", routes.DeleteUserController)
		userGroup.PUT("/:id", routes.UpdateUserController)
		userGroup.GET("/:id", routes.GetUserController)
	}

	authGroup := e.Group("/auth")
	{
		authGroup.GET("/login", routes.Login)
		authGroup.POST("/logout", routes.Logout)
	}

	e.Logger.Fatal(e.Start(":8080"))
	//
}
