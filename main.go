package main

import (
	"i-sports/controller"
	"i-sports/util"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

var engine *xorm.Engine

func main() {
	util.InitEngine()
	initRouter()
}

func initRouter() {
	e := echo.New()
	g := e.Group("/api/v1")
	ctrlProfile := controller.NewProfileController()
	ctrlMeasure := controller.NewMeasurementController()
	g.POST("/login", ctrlProfile.Login)
	g.GET("/profile/:id", ctrlProfile.Profile)
	g.GET("/measure/:userId", ctrlMeasure.Measurement)
	e.Logger.Fatal(e.Start(":3000"))
}

func test(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>API is working 😃</h1>")
}
