package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"microuser/config/database"
	"microuser/models"
	"net/http"
	"strconv"
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

func UpdateUserController(c echo.Context) (err error) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return
	}

	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}

	database.Gorm.Model(models.User{Id: uint(id)}).Updates(&user)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, "true")
}

func DeleteUserController(c echo.Context) (err error) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return
	}

	var user models.User
	database.Gorm.Delete(&user, id)
	if user.Id != 0 {
		return c.JSON(http.StatusExpectationFailed, "Failed to delete a user.")
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserController(c echo.Context) (err error) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return
	}
	//
	var user models.User
	database.Gorm.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

//func (c echo.Context) error {
//	params := make(map[string]string)
//	_ := c.Bind(&params)
//	return c.JSON(http.StatusOK, params["name"])
//}
