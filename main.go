package main

import (
	"github.com/oarielg/BattleGo/controller"
	"github.com/oarielg/BattleGo/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = template.NewTemplate()

	e.GET("/", controller.BattleHandler)
	e.POST("/", controller.BattleHandler)

	e.Logger.Fatal(e.Start(":8090"))
}
