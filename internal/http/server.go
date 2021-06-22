package http

import "github.com/labstack/echo/v4"

func Run() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}
