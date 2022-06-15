package routes

import "github.com/labstack/echo/v4"

func Login(c echo.Context) error {
	return c.String(200, "로그인 성공 입니다.")
}

func Logout(c echo.Context) error {
	return c.String(204, "로그아웃 성공 입니다.")
}
