package main

import (
	"GoFastAfter50/entities"
	"GoFastAfter50/models"
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
	var mod models.Index
	// mod.Messages.Append() = entities.MessageList{}
	//mod.Messages = make([]entities.Message, 4)

	checkData := CheckUserData()

	if checkData {
		mod.Messages = append(mod.Messages, entities.Message{Msg: "Please check user data!"})
	}
	return c.Render(http.StatusOK, "index", mod)
	//return c.String(http.StatusOK, "Hello, World!")
}
