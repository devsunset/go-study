package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.GET("/attachment", func(c echo.Context) error {
		return c.Attachment("attachment.txt", "attachment.txt")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
