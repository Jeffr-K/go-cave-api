package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"microuser/config/database"
	"microuser/models"
	"net/http"
)

func CreateUserController(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}

	result := database.Gorm.Create(&user)
	fmt.Println(result)

	return c.JSON(http.StatusOK, "true")
}

func UpdateUserController(c echo.Context) error {
	return nil
}

func DeleteUserController(c echo.Context) error {
	return nil
}

func GetUserController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
