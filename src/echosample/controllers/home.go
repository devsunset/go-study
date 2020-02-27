package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
}

func (c HomeController) Init(g *echo.Group) {
	g.GET("", c.Get)
}
func (HomeController) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
