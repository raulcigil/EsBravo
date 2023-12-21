package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo) {
	e.Static("/", "")
	//e.Static("/css", "css")
	//e.Static("/js", "js")

	e.GET("/", GetRoot)
}

func GetRoot(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
	//return c.String(http.StatusOK, "Hello, World!")
}
